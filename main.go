// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)

	spew.Dump(r.Header)
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Welcome to Google CloudRun!!")
// 	ctx := context.Background()
// 	if err := run(ctx, os.Stdout, os.Args); err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err)
// 		os.Exit(1)
// 	}
// }

// func run(ctx context.Context, w io.Writer, args []string) error {
// 	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
// 	defer cancel()

// 	// ...
// 	return nil
// }

// // Future ...
// // func run(
// // 	ctx    context.Context,
// // 	args   []string,
// // 	getenv func(string) string,
// // 	stdin  io.Reader,
// // 	stdout, stderr io.Writer,
// // ) error

// func NewServer(
// 	logger *Logger
// 	config *Config
// 	commentStore *commentStore
// 	anotherStore *anotherStore
// ) http.Handler {
// 	mux := http.NewServeMux()
// 	addRoutes(
// 		mux,
// 		Logger,
// 		Config,
// 		commentStore,
// 		anotherStore,
// 	)
// 	var handler http.Handler = mux
// 	handler = someMiddleware(handler)
// 	handler = someMiddleware2(handler)
// 	handler = someMiddleware3(handler)
// 	return handler
// }
