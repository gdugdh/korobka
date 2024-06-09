package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "golang-boilerplate/pkg/chat_v1" // Update
	// gw "github.com/yourorg/yourrepo/proto/gen/go/your/service/v1/your_service"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func CustomMatcher(key string) (string, bool) {
	return key, true
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(CustomMatcher))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterChatV1HandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// mux := http.NewServeMux()
	// mux.HandlePath("GET", "/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// // cors.Default() setup the middleware with default options being
	// // all origins accepted with simple methods (GET, POST). See
	// // documentation below for more options.
	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe(":8080", handler)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
