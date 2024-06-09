package club

import (
	"context"
	"golang-boilerplate/internal/client/db"
	"golang-boilerplate/internal/cors"

	// "golang-boilerplate/internal/model_db"
	"golang-boilerplate/internal/repository"
	desc "golang-boilerplate/pkg/club_v1"
)

type Implementation struct {
	desc.UnimplementedClubV1Server

	clubRepository repository.ClubRepository
	txManager      db.TxManager
}

func NewImplementation(clubRepository repository.ClubRepository, txManager db.TxManager) *Implementation {
	return &Implementation{
		clubRepository: clubRepository,
		txManager:      txManager,
	}
}

func (i *Implementation) ActiveMatchList(ctx context.Context, req *desc.MatchListRequest) (*desc.MatchListResponse, error) {
	cors.ApplyCORSHeaders(ctx)
	matches, err := i.clubRepository.ActiveMatchList(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.MatchListResponse{
		Matches: matches,
	}, nil
}

func (i *Implementation) FinishedMatchList(ctx context.Context, req *desc.FinishedMatchListRequest) (*desc.FinishedMatchListResponse, error) {
	cors.ApplyCORSHeaders(ctx)
	matches, err := i.clubRepository.FinishedMatchList(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.FinishedMatchListResponse{
		Matches: matches,
	}, nil
}

func (i *Implementation) FriendList(ctx context.Context, req *desc.FriendListRequest) (*desc.FriendListResponse, error) {
	cors.ApplyCORSHeaders(ctx)
	friends, err := i.clubRepository.FriendList(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.FriendListResponse{
		Users: friends,
	}, nil
}
