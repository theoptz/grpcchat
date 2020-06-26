package server

import "sync"

import "github.com/theoptz/grpcchat/pkg/chat"

// Handler interface
type Handler interface {
	Add(string, string) bool
	Remove(string) (string, bool)
	GetNameByToken(string) (string, bool)
	AddStream(string) chan chat.StreamResponse
	CloseStream(string)
	Broadcast(chat.StreamResponse)
}

// Storage impements Handler interface
type Storage struct {
	names map[string]map[string]struct{}
	nmu   sync.RWMutex

	tokens map[string]string
	tmu    sync.RWMutex

	streams map[string]chan chat.StreamResponse
	smu     sync.RWMutex
}

// Add method adds client by provided name and token
func (s *Storage) Add(name, token string) bool {
	s.addToken(name, token)

	return s.addName(name, token)
}

// Remove method removes client by provided token if client found
func (s *Storage) Remove(token string) (string, bool) {
	name, ok := s.removeToken(token)
	if !ok {
		return "", false
	}

	return name, s.removeName(name, token)
}

// GetNameByToken returns client name by provided token
func (s *Storage) GetNameByToken(token string) (string, bool) {
	s.tmu.RLock()
	name, ok := s.tokens[token]
	s.tmu.RUnlock()

	return name, ok
}

// AddStream creates new Stream if not exists and associates this stream with provided token
func (s *Storage) AddStream(token string) chan chat.StreamResponse {
	var stream chan chat.StreamResponse
	var ok bool

	s.smu.Lock()
	defer s.smu.Unlock()

	stream, ok = s.streams[token]
	if ok {
		return stream
	}

	stream = make(chan chat.StreamResponse, 100)
	s.streams[token] = stream

	return stream
}

// CloseStream closes stream by provided token
func (s *Storage) CloseStream(token string) {
	s.smu.Lock()
	stream, ok := s.streams[token]
	if ok {
		delete(s.streams, token)
		close(stream)
	}
	s.smu.Unlock()
}

// Broadcast serves for message broadcasting to clients
func (s *Storage) Broadcast(stream chat.StreamResponse) {
	s.smu.RLock()
	for _, clientStream := range s.streams {
		select {
		case clientStream <- stream:
			// ok
		default:
			// client stream is full, dropping message
		}
	}
	s.smu.RUnlock()
}

func (s *Storage) addName(name, token string) bool {
	s.nmu.Lock()
	_, ok := s.names[name]
	if !ok {
		s.names[name] = map[string]struct{}{
			token: struct{}{},
		}
	} else {
		s.names[name][token] = struct{}{}
	}
	s.nmu.Unlock()

	return !ok
}

func (s *Storage) removeName(name, token string) bool {
	s.nmu.Lock()
	defer s.nmu.Unlock()

	_, ok := s.names[name][token]
	if !ok {
		return false
	}

	delete(s.names[name], token)
	if len(s.names[name]) > 0 {
		return false
	}

	delete(s.names, name)
	return true
}

func (s *Storage) addToken(name, token string) {
	s.tmu.Lock()
	s.tokens[token] = name
	s.tmu.Unlock()
}

func (s *Storage) removeToken(token string) (string, bool) {
	s.tmu.Lock()
	name, ok := s.tokens[token]
	if ok {
		delete(s.tokens, token)
	}
	s.tmu.Unlock()

	return name, ok
}

// NewStorage returns Storage pointer
func NewStorage() *Storage {
	return &Storage{
		names:   map[string]map[string]struct{}{},
		tokens:  map[string]string{},
		streams: map[string]chan chat.StreamResponse{},
	}
}
