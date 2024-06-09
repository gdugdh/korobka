package chat

import (
	"context"

	"golang-boilerplate/internal/client/db"
	"golang-boilerplate/internal/model_db"
	"golang-boilerplate/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, message *model_db.Message) error {
	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: "INSERT INTO public.\"Message\" (id_author, id_chat, content, datetime) VALUES ($1, $2, $3, $4);",
	}

	_, err := r.db.DB().ExecContext(ctx, q, message.IdAuthor, message.IdChat, message.Content, message.Datetime)
	return err
}

func (r *repo) GetUserInChat(ctx context.Context, idChat int64) ([]*model_db.User, error) {
	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: "SELECT id, username, email, full_name, avatar FROM public.\"User\" Where id IN (SELECT id FROM public.\"ChatUser\" WHERE id_chat=$1);",
	}

	var users []*model_db.User
	err := r.db.DB().ScanAllContext(ctx, &users, q, idChat)
	return users, err
}

func (r *repo) GetChatMessages(ctx context.Context, idChat int64) ([]*model_db.Message, error) {
	q := db.Query{
		Name:     "chat_repository.GetChatMessages",
		QueryRaw: "SELECT * From public.\"Message\" Where id_chat=$1;",
	}

	var messages []*model_db.Message
	err := r.db.DB().ScanAllContext(ctx, &messages, q, idChat)
	return messages, err
}

// func (r *repo) GetUnreadMessages(ctx context.Context, idUser int64) ([]*model.Message, error) {
// 	q := db.Query{
// 		Name:     "chat_repository.GetUnreadMessages",
// 		QueryRaw: "SELECT * FROM public.\"Message\" WHERE id = $1;",
// 	}

// 	var chat modelRepo.chat
// 	err := r.db.DB().ScanOneContext(ctx, &chat, q, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return converter.TochatFromRepo(&chat), nil
// }
