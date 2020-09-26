package main

import (
	"context"
	"log"
	"net/http"

	"github.com/golangast/collab/go/Contextor"
	Handler "github.com/golangast/collab/go/Handlers/Home"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", Handler.Home)
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))
}
