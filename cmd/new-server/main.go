package main

// import (
// 	"context"
// 	"log"
// 	"strings"

// 	"flag"
// 	"golang-boilerplate/internal/app"
// 	"net/http"

// 	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/grpclog"

// 	// Update
// 	// Update
// 	gw2 "golang-boilerplate/pkg/chat_v1" // Update
// 	gw3 "golang-boilerplate/pkg/club_v1" // Update
// )

// var (
// 	// command-line options:
// 	// gRPC server endpoint
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
// )

// // Options is a set of options to be passed to Run
// type Options struct {
// 	// Addr is the address to listen
// 	Addr string

// 	// Mux is a list of options to be passed to the gRPC-Gateway multiplexer
// 	Mux []gwruntime.ServeMuxOption
// }

// func CustomMatcher(s string) (string, bool) {
// 	return s, true
// }

// func allowCORS(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if origin := r.Header.Get("Origin"); origin != "" {
// 			w.Header().Set("Access-Control-Allow-Origin", origin)
// 			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
// 				preflightHandler(w, r)
// 				return
// 			}
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }

// // preflightHandler adds the necessary headers in order to serve
// // CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// // We insist, don't do this without consideration in production systems.
// func preflightHandler(w http.ResponseWriter, r *http.Request) {
// 	headers := []string{"Content-Type", "Accept", "Authorization"}
// 	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
// 	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
// 	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
// 	log.Printf("Preflight request for %s", r.URL.Path)
// }

// func run(ctx context.Context) error {
// 	opts := Options{}
// 	opts.Addr = "localhost:50051"

// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	conn, err := grpc.NewClient(opts.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return err
// 	}
// 	go func() {
// 		<-ctx.Done()
// 		if err := conn.Close(); err != nil {
// 			log.Fatalf("Failed to close a client connection to the gRPC server: %v", err)
// 		}
// 	}()

// 	mux := http.NewServeMux()

// 	gw := gwruntime.NewServeMux(opts.Mux...)
// 	err = gw2.RegisterChatV1HandlerFromEndpoint(ctx, gw, conn)
// 	if err != nil {
// 		return err
// 	}
// 	err = gw3.RegisterClubV1HandlerFromEndpoint(ctx, gw, conn)
// 	if err != nil {
// 		return err
// 	}
// 	mux.Handle("/", gw)

// 	s := &http.Server{
// 		Addr:    opts.Addr,
// 		Handler: allowCORS(mux),
// 	}

// 	go func() {
// 		<-ctx.Done()
// 		log.Println("Shutting down the http server")
// 		if err := s.Shutdown(context.Background()); err != nil {
// 			log.Fatalf("Failed to shutdown http server: %v", err)
// 		}
// 	}()

// 	log.Printf("Starting listening at %s", opts.Addr)
// 	if err := s.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Fatalf("Failed to listen and serve: %v", err)
// 		return err
// 	}
// 	return nil
// }

// func HttpRunServer() {
// 	flag.Parse()

// 	if err := run(); err != nil {
// 		grpclog.Fatal(err)
// 	}
// }

// func main() {
// 	ctx := context.Background()

// 	a, err := app.NewApp(ctx)
// 	if err != nil {
// 		log.Fatalf("failed to init app: %s", err.Error())
// 	}

// 	go HttpRunServer()

// 	err = a.Run()
// 	if err != nil {
// 		log.Fatalf("failed to run app: %s", err.Error())
// 	}
// }
