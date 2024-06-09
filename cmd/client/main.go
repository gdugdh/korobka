package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"errors"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "golang-boilerplate/pkg/chat_v1"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := desc.NewChatV1Client(conn)

	log.Println("user id:")
	var IdUser int64
	fmt.Scan(&IdUser)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		err = connectChat(ctx, client, IdUser) // , 7*time.Second
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}
	}()

	wg.Wait()
}

func connectChat(ctx context.Context, client desc.ChatV1Client, IdUser int64) error { //, period time.Duration
	stream, err := client.ConnectUser(ctx, &desc.ConnectUserRequest{
		Id: IdUser,
	})
	if err != nil {
		return err
	}

	errch := make(chan error)
	go ReceivingMessage(ctx, stream, IdUser, errch)
	go SendingMessage(ctx, stream, client, IdUser, errch)
	return <-errch
}

func ReceivingMessage(ctx context.Context, stream desc.ChatV1_ConnectUserClient, IdUser int64, errch chan error) {
	for {
		message, errRecv := stream.Recv()
		if errRecv == io.EOF {
			errch <- errors.New("close channel")
			return
		}
		if errRecv != nil {
			errch <- errors.New("failed to receive message from stream")
			return
		}

		if message.IdAuthor != IdUser {
			log.Printf("[%v] - [from: %v]: %s\n", color.YellowString(message.Datetime.AsTime().Format(time.RFC3339)), message.IdAuthor, message.Content)
		}
	}
}

func SendingMessage(ctx context.Context, stream desc.ChatV1_ConnectUserClient, client desc.ChatV1Client, IdUser int64, errch chan error) {
	for {
		var text string
		fmt.Scan(&text)

		_, err := client.SendMessage(ctx, &desc.SendMessageRequest{
			Message: &desc.Message{
				Id:       1,
				IdAuthor: IdUser,
				IdChat:   1,
				Content:  text,
				Datetime: timestamppb.Now(),
			},
		})

		if err != nil {
			errch <- errors.New("failed to send message")
			return
		}
	}
}
