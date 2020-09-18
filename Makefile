# https://github.com/duckladydinh/gomessenger

# fork: github.com/joe-getcouragenow/gomessenger
# WOrks
# Chat App with GRPC go backend demo.

# git include
BOILERPLATE_FSPATH=./../packages/boilerplate
include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/gitr.mk


LIB=github.com/joe-getcouragenow/gomessenger
LIB_FSPATH=$(GOPATH)/src/$(LIB)
LIB_BIN_FSPATH=$(PWD)/bin
LIB_DATA_FSPATH=$(PWD)/data
LIB_BIN=$(LIB_BIN_FSPATH)/gomessenger-$(GO_OS)-$(GO_ARCH)

print:
	@echo
	@echo GO_OS : $(GO_OS)
	@echo GO_ARCH : $(GO_ARCH)
	@echo GIT_VERSION : $(GIT_VERSION)
	@echo
	@echo
	@echo LIB : $(LIB)
	@echo LIB_FSPATH : $(LIB_FSPATH)
	@echo LIB_BIN_FSPATH : $(LIB_BIN_FSPATH)
	@echo LIB_DATA_FSPATH : $(LIB_DATA_FSPATH)
	@echo LIB_BIN : $(LIB_BIN)

dep:
	git clone https://$(LIB).git $(LIB_FSPATH)
dep-delete:
	rm -rf $(LIB_FSPATH)
vscode-add:
	code --add $(LIB_FSPATH) --reuse-window

flu-print:
	@echo flutter devices: 
	flutter devices




### GRPC

evans-dep:
	#go get github.com/ktr0731/evans
	# OR
	brew tap ktr0731/evans
	brew install evans


export PROTO_SOURCE_DIR="proto"
export PROTO_SOURCE_FILE="chat_service.proto"
grpc-gen:
	@echo
	# NOT working !
	#cd $(LIB_FSPATH) && protoc -Iproto proto/chat_service.proto \
		--go_out=plugins=grpc:gms-go/rpc \
		--dart_out=grpc:gms-flutter/lib/rpc
	@echo
	#cd $(LIB_FSPATH) && protoc ${PROTO_SOURCE_DIR}/${PROTO_SOURCE_FILE} --go-grpc_out=proto --go_out=proto --dart_out=proto
	@echo
	#cd $(LIB_FSPATH) && protoc -Iproto proto/chat_service.proto \
		--go_out=gms-go/rpc \
		--go-grpc_out=gms-go/rpc \
		--dart_out=grpc:gms-flutter/lib/rpc
	cd $(LIB_FSPATH) && protoc -Iproto proto/chat_service.proto \
		--go_out=gms-go/rpc \
		--go-grpc_out=gms-go/rpc \
		--dart_out=grpc:gms-flutter/lib/rpc
	@echo

grpc-evans:
	# CLI for GRPC
	# Must have GRPC Reflection turns on in code.
	evans repl $(LIB_FSPATH)/proto/chat_service.proto

#  --js_out=import_style=commonjs:gms-web/src/rpc/ \
#  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:gms-web/src/rpc/ \



### Server
# :9090



go-gen-dep:
	# Install protoc
	#brew install protobuf

	# Install protoc-gen-go
	go get github.com/golang/protobuf/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	stat ${GOPATH}/bin/protoc-gen-go

	# Install protoc-gen-go-grpc
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	stat ${GOPATH}/bin/protoc-gen-go-grpc


	# then make go-gen, to gen the grpc code.

	# then make go-mod-tidy, to ensure the imports that the generated code uses are in te go.mod

go-gen:
	# gen runs from the go folder so that the same go.mod is used for gen and run.
	cd $(LIB_FSPATH)/gms-go && go generate
	# copy to right place hack
	mv $(LIB_FSPATH)/proto/*.go $(LIB_FSPATH)/gms-go/rpc/
	

go-mod-update:
	# force pulling modes
	cd $(LIB_FSPATH)/gms-go && go get -v -t -d ./...

go-mod-tidy:
	cd $(LIB_FSPATH)/gms-go && go mod tidy
	cd $(LIB_FSPATH)/gms-go && go mod verify
	cd $(LIB_FSPATH)/gms-go && go mod download

go-server-buildrun:
	cd $(LIB_FSPATH)/gms-go/main && go build -o $(LIB_BIN)
	
	# run
	$(LIB_BIN)

go-server-run:
	cd $(LIB_FSPATH)/gms-go/main && go run .



### Flutter

## Web tests GRPCWeb works
flu-web-build:
	cd $(LIB_FSPATH)/gms-flutter && flutter build web

flu-web-run:
	cd $(LIB_FSPATH)/gms-flutter && flutter run -d chrome

## DESK tests GRPC works
flu-desk-init:
	cd $(LIB_FSPATH)/gms-flutter && hover init
flu-desk-run:
	cd $(LIB_FSPATH)/gms-flutter && hover run

flu-mob-run:
	cd $(LIB_FSPATH)/gms-flutter && flutter run -d 