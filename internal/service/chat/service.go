package chat

import (
	"context"
	"golang-boilerplate/internal/client/db"
	"golang-boilerplate/internal/converter"
	"golang-boilerplate/internal/model_db"
	"golang-boilerplate/internal/repository"
	"golang-boilerplate/internal/service"
	desc "golang-boilerplate/pkg/chat_v1"
	"log"
	"sync"

	"google.golang.org/grpc/metadata"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager

	users map[int64]desc.ChatV1_ConnectUserServer
	mx    sync.RWMutex
}

func NewService(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		users:          make(map[int64]desc.ChatV1_ConnectUserServer),
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

func (s *serv) ConnectUser(idUser int64, stream desc.ChatV1_ConnectUserServer) error {
	log.Printf("User %v connect\n", idUser)

	s.mx.Lock()
	// Пока с пользователем поддерживается соединение, мы храним соединение с ним в мапе
	s.users[idUser] = stream
	s.mx.Unlock()

	// Здесь горутина прекращает свою работу, пока Context не закроется,
	// что будет сигнализировать о дисконекте пользователя
	<-stream.Context().Done()
	s.mx.Lock()
	// Сразу после дисконекта, мы удаляем пользователя из мапы
	delete(s.users, idUser)
	s.mx.Unlock()
	log.Printf("User %v disconnect\n", idUser)
	return nil
}

func (s *serv) SendMessage(ctx context.Context, message *model_db.Message) error {
	// Аутентифкация
	// log.Printf("\nUSERS CONNECTION: %v\n\n", s.users)

	err := s.chatRepository.Create(ctx, message)
	if err != nil {
		return err
	}

	users, err := s.chatRepository.GetUserInChat(ctx, message.IdChat)
	if err != nil {
		return err
	}

	for _, user := range users {
		s.mx.Lock()
		stream, ok := s.users[user.Id]
		if ok {
			md := metadata.New(map[string]string{"Content-Type": "text/event-stream", "Connection": "keep-alive", "Cache-Control": "no-cache", "X-Accel-Buffering": "no"})
			// srv.SetHeader(md)
			// header := metadata.Pairs(
			// 	"Access-Control-Allow-Origin", "*",
			// )
			stream.SetHeader(md)
			err = stream.Send(converter.MessageM2D(message))
			if err != nil {
				delete(s.users, user.Id)
			}
		}
		s.mx.Unlock()
	}

	return nil
}

func (s *serv) GetChatMessages(ctx context.Context, idChat int64) ([]*model_db.Message, error) {
	messages, err := s.chatRepository.GetChatMessages(ctx, idChat)
	return messages, err
}
