package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/joe-getcouragenow/gomessenger/constants"
	"github.com/joe-getcouragenow/gomessenger/rpc"
	"github.com/joe-getcouragenow/gomessenger/service"
)

func main() {
	server := grpc.NewServer()
	webServer := grpcweb.WrapServer(server, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}))

	// ./main.go:27:32: cannot use chatServiceServer (type *service.ChatServiceServer) as type *rpc.ChatServiceService in argument to rpc.RegisterChatServiceService

	chatServiceServer := service.NewChatServiceServer()
	//chatServiceServer := service.NewChatServiceServer()

	rpc.RegisterChatServiceService(server, &rpc.ChatServiceService{
		RegisterUser: chatServiceServer.RegisterUser,
		AddChatMessage: chatServiceServer.AddChatMessage,
		GetChatMessageStream: chatServiceServer.GetChatMessageStream,
	})

	httpServer := http.Server{
		Addr: constants.ServerAddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			webServer.ServeHTTP(w, r)
		}),
	}

	log.Println("Starting server at", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
