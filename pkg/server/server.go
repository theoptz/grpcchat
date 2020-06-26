package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/theoptz/grpcchat/internal/debug"
	"github.com/theoptz/grpcchat/internal/utils"
	"github.com/theoptz/grpcchat/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Server struct
type Server struct {
	addr          string
	logger        debug.Logger
	clients       Handler
	broadcastChan chan chat.StreamResponse
}

// New returns Server pointer
func New(addr string, isDebug bool) (*Server, error) {
	if addr == "" {
		return nil, errors.New("Invalid address")
	}

	return &Server{
		addr:          addr,
		logger:        debug.New(isDebug),
		clients:       NewStorage(),
		broadcastChan: make(chan chat.StreamResponse, 1000),
	}, nil
}

// Run starts server
func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	chat.RegisterChatServer(srv, s)

	s.logger.Debug("Server listening on %s", s.addr)

	go s.broadcast(ctx)

	go func() {
		err := srv.Serve(lis)
		if err != nil {
			log.Println("Serve error", err)
		}

		cancel()
	}()

	<-ctx.Done()
	srv.GracefulStop()
	s.logger.Debug("Server has been shutdown")

	return nil
}

// Login handles client login request
func (s *Server) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid name")
	}

	token := utils.GenerateToken(req.Name)
	ok := s.clients.Add(req.Name, token)

	s.logger.Debug("%s connected to server with token %s", req.Name, token)

	if ok {
		s.logger.Debug("%s goes online", req.Name)

		s.broadcastChan <- chat.StreamResponse{
			Timestamp: ptypes.TimestampNow(),
			Event: &chat.StreamResponse_UserLogin{
				UserLogin: &chat.StreamResponse_Login{
					Name: req.Name,
				},
			},
		}
	}

	return &chat.LoginResponse{
		Token: token,
	}, nil
}

// Logout handles client logout request
func (s *Server) Logout(ctx context.Context, req *chat.LogoutRequest) (*chat.LogoutResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid token")
	}

	name, ok := s.clients.Remove(req.Token)
	if !ok {
		if name == "" {
			return nil, status.Error(codes.NotFound, "Token not found")
		}

		s.logger.Debug("%s (%s) disconnected from server", name, req.Token)
		return new(chat.LogoutResponse), nil
	}

	s.logger.Debug("%s (%s) disconnected from server", name, req.Token)
	s.logger.Debug("%s goes offline", name)

	s.broadcastChan <- chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_UserLogout{
			UserLogout: &chat.StreamResponse_Logout{
				Name: name,
			},
		},
	}

	return new(chat.LogoutResponse), nil
}

// Stream provides incomming/outcomming message handlers
func (s *Server) Stream(srv chat.Chat_StreamServer) error {
	token, ok := s.getToken(srv.Context())
	if !ok || token == "" {
		return status.Error(codes.Unauthenticated, "Missing token header")
	}

	name, ok := s.clients.GetNameByToken(token)
	if !ok {
		return status.Error(codes.Unauthenticated, "Invalid token")
	}

	go s.sendEventsToClient(srv, token)

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		s.logger.Debug("%s (%s) has sent a message: %s", name, token, req.Message)

		s.broadcastChan <- chat.StreamResponse{
			Timestamp: ptypes.TimestampNow(),
			Event: &chat.StreamResponse_UserMessage{
				UserMessage: &chat.StreamResponse_Message{
					Name:    name,
					Message: req.Message,
				},
			},
		}
	}

	<-srv.Context().Done()
	return srv.Context().Err()
}

func (s *Server) sendEventsToClient(srv chat.Chat_StreamServer, token string) {
	stream := s.clients.AddStream(token)
	defer s.clients.CloseStream(token)

	for {
		select {
		case <-srv.Context().Done():
			// server terminated
			return

		// read new event
		case res := <-stream:
			if r, ok := status.FromError(srv.Send(&res)); ok {
				switch r.Code() {
				case codes.OK:
					// everything's ok

				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					// client terminated
					s.logger.Debug("Client (%s) terminated connection", token)
					return

				default:
					s.logger.Debug("Failed to send to client (%s): %d %v", token, r.Code(), r.Err())
					return
				}
			}
		}
	}
}

func (s *Server) broadcast(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case res := <-s.broadcastChan:
			s.clients.Broadcast(res)
		}
	}
}

func (s *Server) getToken(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[chat.TokenHeader]) == 0 {
		return "", false
	}

	return md[chat.TokenHeader][0], true
}
