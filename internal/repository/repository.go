package repository

import (
	"context"

	"golang-boilerplate/internal/model_db"
	desc "golang-boilerplate/pkg/club_v1"
)

type ClubRepository interface {
	FinishedMatchList(ctx context.Context, id_discipline int64) ([]*desc.Match, error)
	ActiveMatchList(ctx context.Context, id_discipline int64) ([]*desc.Match, error)
	FriendList(ctx context.Context, id_user int64) ([]*desc.User, error)
	// Create(ctx context.Context, club *model_db.Club) (int64, error)
	// Get(ctx context.Context, id int64) (*model_db.Club, error)
	// Update(ctx context.Context, ClubUpdate *model_db.ClubUpdate) error
}

type ChatRepository interface {
	Create(ctx context.Context, message *model_db.Message) error
	GetUserInChat(ctx context.Context, idChat int64) ([]*model_db.User, error)
	GetChatMessages(ctx context.Context, idChat int64) ([]*model_db.Message, error)
	// GetUnreadMessages(ctx context.Context, idUser int64) ([]*model.Message, error)
}
