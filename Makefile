LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

generate-http-note-api:
	protoc -I $GOPATH/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20240603083838-716a2814e199/ -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative api/note_v1/note.proto


# export GOPATH=$HOME/go
# export PATH=$PATH:$GOROOT/bin:$GOPATH/bin 
generate-http-club-api:
	protoc -I $GOPATH/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20240603083838-716a2814e199/ \
	-I . --grpc-gateway_out pkg/club_v1 \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	api/club_v1/club.proto

generate: generate-club-api generate-chat-api

generate-club-api:
	mkdir -p pkg/club_v1
	protoc --proto_path api/club_v1 \
	--go_out=pkg/club_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/club_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/club_v1/club.proto

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

# generate-note-api:
# 	mkdir -p pkg/note_v1
# 	protoc --proto_path api/note_v1 \
# 	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
# 	--plugin=protoc-gen-go=bin/protoc-gen-go \
# 	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
# 	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
# 	api/note_v1/note.proto
