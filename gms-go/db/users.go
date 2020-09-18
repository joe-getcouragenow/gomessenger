package db

import (
	"sync"

	"github.com/duckladydinh/gomessenger/api"
)

// UserDB ...
type UserDB struct {
	data map[string]*api.ChatUser
	mux  sync.Mutex
}

// NewUserDB ..
func NewUserDB() *UserDB {
	return &UserDB{
		data: map[string]*api.ChatUser{},
	}
}

// AddUser ...
func (s *UserDB) AddUser(user *api.ChatUser) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[user.Id] = user
}

// GetUser ...
func (s *UserDB) GetUser(userId string) *api.ChatUser {
	s.mux.Lock()
	defer s.mux.Unlock()

	if v, ok := s.data[userId]; ok {
		return v
	}
	return nil
}
