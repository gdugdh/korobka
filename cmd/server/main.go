package main

import (
	"context"
	"log"
	"strings"

	"flag"
	"golang-boilerplate/internal/app"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "golang-boilerplate/pkg/chat_v1"  // Update
	gw2 "golang-boilerplate/pkg/club_v1" // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

//	func CustomMatcher(s string) (string, bool) {
//		return s, true
//	}
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Printf("Preflight request for %s", r.URL.Path)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux_runtime := runtime.NewServeMux(
	// runtime.WithIncomingHeaderMatcher(CustomMatcher),
	// runtime.WithOutgoingHeaderMatcher(CustomMatcher),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterChatV1HandlerFromEndpoint(ctx, mux_runtime, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = gw2.RegisterClubV1HandlerFromEndpoint(ctx, mux_runtime, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", mux_runtime)

	s := &http.Server{
		Addr:    "localhost:8081",
		Handler: allowCORS(mux),
	}

	log.Println("HTTP Proxy server is running on localhost:8081")
	return s.ListenAndServe()
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	// return http.ListenAndServe(":8081", mux)
}

func HttpRunServer() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	go HttpRunServer()

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
