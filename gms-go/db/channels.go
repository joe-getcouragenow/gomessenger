package db

import (
	"sync"

	"github.com/duckladydinh/gomessenger/api"
	
)
// ChannelDB ...
type ChannelDB struct {
	data map[string]chan *api.ChatMessage
	mux  sync.Mutex
}

// NewChannelDB ...
func NewChannelDB() *ChannelDB {
	return &ChannelDB{
		data: make(map[string]chan *api.ChatMessage),
	}
}

// NewChannel ...
func (s *ChannelDB) NewChannel(userId string) chan *api.ChatMessage {
	s.mux.Lock()
	defer s.mux.Unlock()

	ch := make(chan *api.ChatMessage)
	s.data[userId] = ch
	return ch
}

// RemoveChannel ...
func (s *ChannelDB) RemoveChannel(userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.data, userId)
}

// Broadcast ...
func (s *ChannelDB) Broadcast(msg *api.ChatMessage) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, ch := range s.data {
		ch <- msg
	}
}
