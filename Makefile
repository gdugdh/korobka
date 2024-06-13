LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/server/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get github.com/googleapis/googleapis@v0.0.0-20240603083838-716a2814e199

generate: generate-club-api generate-chat-api

generate-club-api:
	mkdir -p pkg/club_v1
	protoc --proto_path api/club_v1 \
	--go_out=pkg/club_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/club_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/club_v1/club.proto

generate-http-club-api:
	export GOPATH=$(HOME)/go
	export PATH=$(PATH):$(GOROOT)/bin:$(GOPATH)/bin
	protoc -I $(HOME)/go/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20240603083838-716a2814e199/ \
	-I . --grpc-gateway_out pkg/club_v1 \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	api/club_v1/club.proto
	cp pkg/club_v1/api/club_v1/club.pb.gw.go pkg/club_v1/
	rm -rf pkg/club_v1/api

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

generate-http-chat-api:
	export GOPATH=$(HOME)/go
	export PATH=$(PATH):$(GOROOT)/bin:$(GOPATH)/bin
	protoc -I $(HOME)/go/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20240603083838-716a2814e199/ \
	-I . --grpc-gateway_out pkg/chat_v1 \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	api/chat_v1/chat.proto
	cp pkg/chat_v1/api/chat_v1/chat.pb.gw.go pkg/chat_v1/
	rm -rf pkg/chat_v1/api
build:
	go build -o bin/korobka .
