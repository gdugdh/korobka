package app

import (
	"context"
	"log"

	"golang-boilerplate/internal/api/chat"
	"golang-boilerplate/internal/api/club"

	"golang-boilerplate/internal/client/db"
	"golang-boilerplate/internal/client/db/pg"
	"golang-boilerplate/internal/client/db/transaction"
	"golang-boilerplate/internal/closer"
	"golang-boilerplate/internal/config"
	"golang-boilerplate/internal/repository"
	"golang-boilerplate/internal/service"

	chatService "golang-boilerplate/internal/service/chat"

	chatRepository "golang-boilerplate/internal/repository/chat"
	clubRepository "golang-boilerplate/internal/repository/club"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	clubRepository repository.ClubRepository
	chatRepository repository.ChatRepository

	chatService service.ChatService

	clubImpl *club.Implementation
	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// CLUB
func (s *serviceProvider) ClubRepository(ctx context.Context) repository.ClubRepository {
	if s.clubRepository == nil {
		s.clubRepository = clubRepository.NewRepository(s.DBClient(ctx))
	}

	return s.clubRepository
}

func (s *serviceProvider) ClubImpl(ctx context.Context) *club.Implementation {
	if s.clubImpl == nil {
		s.clubImpl = club.NewImplementation(
			s.ClubRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.clubImpl
}

// CHAT
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
