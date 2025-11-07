package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eliastarin/ssp2-southpark/go-api/adapters"
	"github.com/eliastarin/ssp2-southpark/go-api/app"
	"github.com/eliastarin/ssp2-southpark/go-api/ports"
)

func buildPublisher() ports.MessagePublisher {
	amqpURL := os.Getenv("AMQP_URL")
	queue := os.Getenv("AMQP_QUEUE")
	if queue == "" {
		queue = "southpark_messages"
	}

	if amqpURL != "" {
		pub, err := adapters.NewRabbitPublisher(amqpURL, queue)
		if err == nil {
			log.Printf("RabbitMQ publisher connected: %s queue=%s", amqpURL, queue)
			return pub
		}
		log.Printf("RabbitMQ unavailable (%v). Falling back to memory publisher.", err)
	} else {
		log.Printf("AMQP_URL not set; using memory publisher.")
	}

	return adapters.NewMemoryPublisher()
}

func main() {
	pub := buildPublisher()
	h := app.NewHandlers(pub)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/messages", h.PostMessage)

	addr := ":8080"
	log.Printf("go-api listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
