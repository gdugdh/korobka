package service

import (
	"context"

	"golang-boilerplate/internal/model_db"
	desc "golang-boilerplate/pkg/chat_v1"
)

type ChatService interface {
	SendMessage(ctx context.Context, message *model_db.Message) error
	ConnectUser(idUser int64, stream desc.ChatV1_ConnectUserServer) error
	GetChatMessages(ctx context.Context, idChat int64) ([]*model_db.Message, error)
}
