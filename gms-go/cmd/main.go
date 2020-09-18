package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/duckladydinh/gomessenger/constants"
	"github.com/duckladydinh/gomessenger/rpc"
	"github.com/duckladydinh/gomessenger/service"
)

func main() {
	server := grpc.NewServer()
	webServer := grpcweb.WrapServer(server, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}))

	// ./main.go:27:32: cannot use chatServiceServer (type *service.ChatServiceServer) as type *rpc.ChatServiceService in argument to rpc.RegisterChatServiceService

	chatServiceServer := service.NewChatServiceServer()
	//chatServiceServer := service.NewChatServiceServer()

	rpc.RegisterChatServiceService(server, chatServiceServer)
	//rpc.RegisterChatServiceServer(server, chatServiceServer)
	//rpc.RegisterChatServiceService(server, chatServiceServer)

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
