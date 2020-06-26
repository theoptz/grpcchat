package server

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/theoptz/grpcchat/pkg/chat"
)

func TestStorageAddToken(t *testing.T) {
	cases := []struct {
		before map[string]string
		after  map[string]string
		token  string
		name   string
	}{
		{
			before: map[string]string{},
			after: map[string]string{
				"token": "name",
			},
			token: "token",
			name:  "name",
		},
		{
			before: map[string]string{
				"token": "name",
			},
			after: map[string]string{
				"token": "name",
				"foo":   "bar",
			},
			token: "foo",
			name:  "bar",
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.tokens = tc.before

		s.addToken(tc.name, tc.token)

		if !reflect.DeepEqual(s.tokens, tc.after) {
			t.Errorf("Map must be %#v but got %#v", tc.after, s.tokens)
		}
	}
}

func TestStorageRemoveToken(t *testing.T) {
	cases := []struct {
		before map[string]string
		after  map[string]string
		token  string
		name   string
		ok     bool
	}{
		{
			before: map[string]string{
				"token": "name",
			},
			after: map[string]string{},
			token: "token",
			name:  "name",
			ok:    true,
		},
		{
			before: map[string]string{
				"token": "name",
			},
			after: map[string]string{
				"token": "name",
			},
			token: "token2",
			name:  "",
			ok:    false,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.tokens = tc.before

		name, ok := s.removeToken(tc.token)

		if ok != tc.ok {
			t.Errorf("Name must be %s but got %s", tc.name, name)
			continue
		}

		if ok != tc.ok {
			t.Errorf("Res must be %t but got %t", tc.ok, ok)
			continue
		}

		if !reflect.DeepEqual(s.tokens, tc.after) {
			t.Errorf("Map must be %#v but got %#v", tc.after, s.tokens)
		}
	}
}

func TestStorageAddName(t *testing.T) {
	cases := []struct {
		before map[string]map[string]struct{}
		after  map[string]map[string]struct{}
		name   string
		token  string
		ok     bool
	}{
		{
			before: map[string]map[string]struct{}{},
			after: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			name:  "name",
			token: "token",
			ok:    true,
		},
		{
			before: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			after: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token":  struct{}{},
					"token2": struct{}{},
				},
			},
			name:  "name",
			token: "token2",
			ok:    false,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.names = tc.before

		ok := s.addName(tc.name, tc.token)

		if tc.ok != ok {
			t.Errorf("Res must be %t but got %t", tc.ok, ok)
			continue
		}

		if !reflect.DeepEqual(s.names, tc.after) {
			t.Errorf("Map must be %#v but got %#v", tc.after, s.tokens)
		}
	}
}

func TestStorageRemoveName(t *testing.T) {
	cases := []struct {
		before map[string]map[string]struct{}
		after  map[string]map[string]struct{}
		name   string
		token  string
		ok     bool
	}{
		{
			before: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			after: map[string]map[string]struct{}{},
			name:  "name",
			token: "token",
			ok:    true,
		},
		{
			before: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			after: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			name:  "name",
			token: "token2",
			ok:    false,
		},
		{
			before: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token":  struct{}{},
					"token2": struct{}{},
				},
			},
			after: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token2": struct{}{},
				},
			},
			name:  "name",
			token: "token",
			ok:    false,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.names = tc.before

		ok := s.removeName(tc.name, tc.token)

		if tc.ok != ok {
			t.Errorf("Res must be %t but got %t", tc.ok, ok)
			continue
		}

		if !reflect.DeepEqual(s.names, tc.after) {
			t.Errorf("Map must be %#v but got %#v", tc.after, s.tokens)
		}
	}
}

func TestStorageGetNameByToken(t *testing.T) {
	cases := []struct {
		tokens map[string]string
		token  string
		name   string
		ok     bool
	}{
		{
			tokens: map[string]string{},
			token:  "example",
			name:   "",
			ok:     false,
		},
		{
			tokens: map[string]string{
				"token": "name",
			},
			token: "token",
			name:  "name",
			ok:    true,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.tokens = tc.tokens

		name, ok := s.GetNameByToken(tc.token)

		if tc.name != name {
			t.Errorf("Name must be %s but got %s", tc.name, name)
			continue
		}

		if tc.ok != ok {
			t.Errorf("Ok must be %t but got %t", tc.ok, ok)
		}
	}
}

func TestStorageAdd(t *testing.T) {
	cases := []struct {
		tokensBefore map[string]string
		tokensAfter  map[string]string
		namesBefore  map[string]map[string]struct{}
		namesAfter   map[string]map[string]struct{}
		token        string
		name         string
		res          bool
	}{
		{
			tokensBefore: map[string]string{},
			tokensAfter: map[string]string{
				"token": "name",
			},
			namesBefore: map[string]map[string]struct{}{},
			namesAfter: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			name:  "name",
			token: "token",
			res:   true,
		},
		{
			tokensBefore: map[string]string{
				"token": "name",
			},
			tokensAfter: map[string]string{
				"token":  "name",
				"token2": "name",
			},
			namesBefore: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			namesAfter: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token":  struct{}{},
					"token2": struct{}{},
				},
			},
			name:  "name",
			token: "token2",
			res:   false,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.names = tc.namesBefore
		s.tokens = tc.tokensBefore

		res := s.Add(tc.name, tc.token)

		if res != tc.res {
			t.Errorf("Res must be %t but got %t", tc.res, res)
			continue
		}

		if !reflect.DeepEqual(tc.tokensAfter, s.tokens) {
			t.Errorf("Tokens must be %#v but got %#v", tc.tokensAfter, s.tokens)
			continue
		}

		if !reflect.DeepEqual(tc.namesAfter, s.names) {
			t.Errorf("Names must be %#v but got %#v", tc.namesAfter, s.names)
		}
	}
}

func TestStorageRemove(t *testing.T) {
	cases := []struct {
		tokensBefore map[string]string
		tokensAfter  map[string]string
		namesBefore  map[string]map[string]struct{}
		namesAfter   map[string]map[string]struct{}
		token        string
		name         string
		ok           bool
	}{
		{
			tokensBefore: map[string]string{
				"token": "name",
			},
			tokensAfter: map[string]string{},
			namesBefore: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			namesAfter: map[string]map[string]struct{}{},
			token:      "token",
			name:       "name",
			ok:         true,
		},
		{
			tokensBefore: map[string]string{},
			tokensAfter:  map[string]string{},
			namesBefore:  map[string]map[string]struct{}{},
			namesAfter:   map[string]map[string]struct{}{},
			token:        "token",
			name:         "",
			ok:           false,
		},
		{
			tokensBefore: map[string]string{
				"token":  "name",
				"token2": "name",
			},
			tokensAfter: map[string]string{
				"token": "name",
			},
			namesBefore: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token":  struct{}{},
					"token2": struct{}{},
				},
			},
			namesAfter: map[string]map[string]struct{}{
				"name": map[string]struct{}{
					"token": struct{}{},
				},
			},
			token: "token2",
			name:  "name",
			ok:    false,
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.names = tc.namesBefore
		s.tokens = tc.tokensBefore

		name, ok := s.Remove(tc.token)

		if name != tc.name {
			t.Errorf("Name must be %s but got %s", tc.name, name)
			continue
		}

		if ok != tc.ok {
			t.Errorf("Ok must be %t but got %t", tc.ok, ok)
			continue
		}

		if !reflect.DeepEqual(tc.tokensAfter, s.tokens) {
			t.Errorf("Tokens must be %#v but got %#v", tc.tokensAfter, s.tokens)
			continue
		}

		if !reflect.DeepEqual(tc.namesAfter, s.names) {
			t.Errorf("Names must be %#v but got %#v", tc.namesAfter, s.names)
		}
	}
}

func TestStorageAddStream(t *testing.T) {
	cases := []struct {
		streams map[string]chan chat.StreamResponse
		token   string
	}{
		{
			streams: map[string]chan chat.StreamResponse{},
			token:   "example",
		},
		{
			streams: map[string]chan chat.StreamResponse{
				"example": make(chan chat.StreamResponse),
			},
			token: "example",
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.streams = tc.streams

		stream := s.AddStream(tc.token)

		if str, _ := s.streams[tc.token]; !reflect.DeepEqual(stream, str) {
			t.Errorf("Must return stream %#v from map but got %#v", str, stream)
		}
	}
}

func TestStorageRemoveStream(t *testing.T) {
	stream := make(chan chat.StreamResponse)

	cases := []struct {
		before map[string]chan chat.StreamResponse
		after  map[string]chan chat.StreamResponse
		token  string
	}{
		{
			before: map[string]chan chat.StreamResponse{},
			after:  map[string]chan chat.StreamResponse{},
			token:  "example",
		},
		{
			before: map[string]chan chat.StreamResponse{
				"example": stream,
			},
			after: map[string]chan chat.StreamResponse{},
			token: "example",
		},
	}

	for _, tc := range cases {
		s := NewStorage()
		s.streams = tc.before

		s.CloseStream(tc.token)

		if !reflect.DeepEqual(s.streams, tc.after) {
			t.Errorf("Streams must be %#v but got %#v", tc.after, s.streams)
		}
	}
}

func TestStorageBroadcast(t *testing.T) {
	s := NewStorage()

	l := 10
	s.streams = map[string]chan chat.StreamResponse{}

	stream := make(chan chat.StreamResponse, l)

	for i := 0; i < l; i++ {
		s.streams["token"+strconv.Itoa(i)] = stream
	}

	go s.Broadcast(chat.StreamResponse{})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	count := 0

	for {
		select {
		case <-ctx.Done():
			t.Errorf("Must receive %d messages but got only %d", l, count)
			return

		case <-stream:
			count++
			if count == l {
				return
			}
		}
	}
}
