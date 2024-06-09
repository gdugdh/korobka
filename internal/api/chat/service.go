package chat

import (
	"context"
	"golang-boilerplate/internal/service"
	"log"

	"golang-boilerplate/internal/converter"
	desc "golang-boilerplate/pkg/chat_v1"

	"golang-boilerplate/internal/cors"

	"github.com/golang/protobuf/ptypes/empty"
)

// type ConnectedUser struct {
// 	stream desc.ChatV1_ConnectUserServer
// 	// mx     sync.RWMutex
// }

type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}

func (i *Implementation) ConnectUser(req *desc.ConnectUserRequest, stream desc.ChatV1_ConnectUserServer) error {
	// Аутентифкация
	// srv.SetHeader(md)
	// header := metadata.Pairs(
	// 	"Access-Control-Allow-Origin", "*",
	// )
	return i.chatService.ConnectUser(req.GetId(), stream)
}

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	// cors.ApplyCORSHeaders(ctx)
	log.Printf("%v", req.GetMessage())
	err := i.chatService.SendMessage(ctx, converter.MessageD2M(req.GetMessage()))
	if err != nil {
		return nil, err
	}

	log.Printf("sendMessage from: %d", req.Message.GetIdAuthor())

	return &empty.Empty{}, nil
}

func (i *Implementation) GetChatMessages(ctx context.Context, req *desc.GetChatMessagesRequest) (*desc.GetChatMessagesResponse, error) {
	cors.ApplyCORSHeaders(ctx)
	data, err := i.chatService.GetChatMessages(ctx, req.GetId())

	var messages []*desc.Message
	for _, d := range data {
		messages = append(messages, converter.MessageM2D(d))
	}

	return &desc.GetChatMessagesResponse{
		Messages: messages,
	}, err
}
