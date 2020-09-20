package service

import (
	"context"
	"log"
	"sync"

	"github.com/joe-getcouragenow/gomessenger/api"
	"github.com/joe-getcouragenow/gomessenger/db"
	"github.com/joe-getcouragenow/gomessenger/rpc"
	"github.com/joe-getcouragenow/gomessenger/util"
)

// ChatServiceServer ...
type ChatServiceServer struct {
	messageStore *db.MessageDB
	userStore    *db.UserDB
	channelStore *db.ChannelDB
	mux          sync.Mutex
}

// NewChatServiceServer ...
func NewChatServiceServer() *ChatServiceServer {
	return &ChatServiceServer{
		messageStore: db.NewMessageDB(),
		userStore:    db.NewUserDB(),
		channelStore: db.NewChannelDB(),
	}
}

// RegisterUser ...
func (server *ChatServiceServer) RegisterUser(_ context.Context, r *rpc.RegisterUserRequest) (*rpc.RegisterUserResponse, error) {
	server.mux.Lock()
	defer server.mux.Unlock()

	log.Println(".RegisterUser", r)

	userId := util.NewUserId()
	server.userStore.AddUser(&api.ChatUser{
		Id:   userId,
		Name: r.Name,
	})

	res := &rpc.RegisterUserResponse{UserId: userId}
	return res, nil
}

// AddChatMessage ...
func (server *ChatServiceServer) AddChatMessage(_ context.Context, r *rpc.AddChatMessageRequest) (*rpc.Empty, error) {
	server.mux.Lock()
	defer server.mux.Unlock()

	log.Println(".AddChatMessage", r)

	msgId := util.NewMessageId()
	msg := &api.ChatMessage{
		Id:        msgId,
		UserId:    r.UserId,
		Timestamp: r.Timestamp,
		Message:   r.Message,
	}
	server.messageStore.AddMessage(msg)
	server.channelStore.Broadcast(msg)

	return new(rpc.Empty), nil
}

// GetChatMessageStream ...
func (server *ChatServiceServer) GetChatMessageStream(r *rpc.GetChatMessageStreamRequest, s rpc.ChatService_GetChatMessageStreamServer) error {
	server.mux.Lock()

	log.Println(".GetChatChannel", r)

	messages := server.messageStore.GetMessages(0)
	for _, msg := range messages {
		e := s.Send(&rpc.ChatMessage{
			Id:        msg.Id,
			UserId:    msg.UserId,
			Timestamp: msg.Timestamp,
			Message:   msg.Message,
		})
		if e != nil {
			return e
		}
	}
	ch := server.channelStore.NewChannel(r.UserId)

	server.mux.Unlock()

	for msg := range ch {
		e := s.Send(&rpc.ChatMessage{
			Id:        msg.Id,
			UserId:    msg.UserId,
			Timestamp: msg.Timestamp,
			Message:   msg.Message,
		})
		if e != nil {
			server.channelStore.RemoveChannel(r.UserId)
			return e
		}
	}

	return nil
}
