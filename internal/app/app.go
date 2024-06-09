package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"golang-boilerplate/internal/closer"
	"golang-boilerplate/internal/config"
	descChat "golang-boilerplate/pkg/chat_v1"
	descClub "golang-boilerplate/pkg/club_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// type ImplementationDiscipline struct {
// 	descDiscipline.UnimplementedDisciplineV1Server
// }

// type ImplementationUser struct {
// 	descUser.UnimplementedUserV1Server
// }

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	// descNote.RegisterNoteV1Server(a.grpcServer, a.serviceProvider.NoteImpl(ctx))
	descClub.RegisterClubV1Server(a.grpcServer, a.serviceProvider.ClubImpl(ctx))
	descChat.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))
	// descDiscipline.RegisterDisciplineV1Server(a.grpcServer, ImplementationDiscipline{})
	// descUser.RegisterUserV1Server(a.grpcServer, ImplementationUser{})

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
