package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang-boilerplate/internal/model_db"
	desc "golang-boilerplate/pkg/chat_v1"
)

func MessageM2D(message *model_db.Message) *desc.Message {
	return &desc.Message{
		Id:       message.Id,
		IdAuthor: message.IdAuthor,
		IdChat:   message.IdChat,
		Content:  message.Content,
		Datetime: timestamppb.New(message.Datetime),
	}
}

func MessageD2M(message *desc.Message) *model_db.Message {
	return &model_db.Message{
		Id:       message.Id,
		IdAuthor: message.IdAuthor,
		IdChat:   message.IdChat,
		Content:  message.Content,
		Datetime: message.GetDatetime().AsTime(),
	}
}
