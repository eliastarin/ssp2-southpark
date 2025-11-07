package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/eliastarin/ssp2-southpark/go-api/adapters"
	"github.com/eliastarin/ssp2-southpark/go-api/app"
	"github.com/eliastarin/ssp2-southpark/go-api/ports"
)

//go:embed web/*
var webFS embed.FS

func buildPublisher() ports.MessagePublisher {
	amqpURL := os.Getenv("AMQP_URL")
	queue := os.Getenv("AMQP_QUEUE")
	if queue == "" {
		queue = "southpark_messages"
	}
	pub, err := adapters.NewRabbitPublisher(amqpURL, queue)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Printf("RabbitMQ publisher connected: %s queue=%s", amqpURL, queue)
	return pub
}

func main() {
	pub := buildPublisher()
	h := app.NewHandlers(pub)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/messages", h.PostMessage)

	// Web app
	sub, _ := fs.Sub(webFS, "web")
	mux.Handle("/", http.FileServer(http.FS(sub)))

	addr := ":8080"
	log.Printf("go-api listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
