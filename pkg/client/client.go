package client

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/theoptz/grpcchat/internal/debug"
	"github.com/theoptz/grpcchat/pkg/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	connTimeout = time.Duration(5 * time.Second)
)

// Client struct
type Client struct {
	addr string
	name string

	token  string
	logger debug.Logger
	chat   chat.ChatClient

	mu     sync.RWMutex
	closed bool
}

// Run starts a client
func (c *Client) Run(ctx context.Context) error {
	connCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	conn, err := grpc.DialContext(connCtx, c.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.WithMessage(err, "failed to connect to provided address")
	}
	defer conn.Close()

	c.logger.Debug("%s is connected to %s", c.name, c.addr)

	c.chat = chat.NewChatClient(conn)

	c.token, err = c.login(ctx)
	if err != nil {
		return errors.WithMessage(err, "failed to login")
	}

	c.logger.Debug("Logged in successfully as %s", c.name)

	err = c.stream(ctx)
	if err != nil {
		return errors.WithMessage(err, "failed to stream")
	}

	c.logger.Debug("Logging out")
	err = c.logout()
	if err != nil {
		return errors.WithMessage(err, "failed to log out")
	}

	return nil
}

func (c *Client) login(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	res, err := c.chat.Login(ctx, &chat.LoginRequest{
		Name: c.name,
	})
	if err != nil {
		return "", err
	} else if res == nil {
		return "", errors.New("Invalid login response")
	}

	return res.Token, nil
}

func (c *Client) isClosed() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.closed
}

func (c *Client) markAsClosed() {
	c.mu.Lock()
	c.closed = true
	c.mu.Unlock()
}

func (c *Client) logout() error {
	if c.isClosed() {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	_, err := c.chat.Logout(ctx, &chat.LogoutRequest{
		Token: c.token,
	})

	if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
		return nil
	}

	return err
}

func (c *Client) stream(ctx context.Context) error {
	// attach token for outgoing requests
	md := metadata.New(map[string]string{
		chat.TokenHeader: c.token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := c.chat.Stream(ctx)
	if err != nil {
		return err
	}
	defer client.CloseSend()

	c.logger.Debug("Connected to stream")

	// run send/receive methods
	go c.send(client)
	return c.receive(client)
}

func (c *Client) send(client chat.Chat_StreamClient) error {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for {
		select {
		case <-client.Context().Done():
			c.logger.Debug("Client send loop disconnected")
			c.markAsClosed()
			return nil

		default:
			if sc.Scan() {
				if err := client.Send(&chat.StreamRequest{
					Message: sc.Text(),
				}); err != nil {
					c.logger.Debug("Send failed: %v", err)
					return err
				}
			} else {
				err := sc.Err()
				c.logger.Debug("Input scanner failure: %v", err)
				return err
			}
		}
	}
}

func (c *Client) receive(client chat.Chat_StreamClient) error {
	for {
		res, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				c.logger.Debug("Stream closed by server")
				return nil
			}

			s, ok := status.FromError(err)
			if ok {
				code := s.Code()
				if code == codes.Canceled || code == codes.Unavailable {
					c.logger.Debug("Stream canceled (usually indicates shutdown)")
					return nil
				}
			}

			return err
		}

		switch event := res.Event.(type) {
		case *chat.StreamResponse_UserLogin:
			log.Printf("Server: %s is online", event.UserLogin.Name)

		case *chat.StreamResponse_UserLogout:
			log.Printf("Server: %s is offline", event.UserLogout.Name)

		case *chat.StreamResponse_UserMessage:
			log.Printf("%s: %s", event.UserMessage.Name, event.UserMessage.Message)

		case *chat.StreamResponse_ServerShutdown:
			c.logger.Debug("The server is shutting down %#v", event)
			c.markAsClosed()

			return nil

		default:
			c.logger.Debug("Unexpected event from the server: %T", event)

		}
	}
}

// New returns Client pointer
func New(addr, name string, isDebug bool) (*Client, error) {
	if addr == "" {
		return nil, errors.New("Invalid address")
	}

	if name == "" {
		return nil, errors.New("Invalid name")
	}

	return &Client{
		addr:   addr,
		name:   name,
		logger: debug.New(isDebug),
	}, nil
}
