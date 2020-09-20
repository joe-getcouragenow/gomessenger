package main

//go:generate protoc -I./../proto -I. --go_out=paths=source_relative:./rpc ./../proto/chat_service.proto
