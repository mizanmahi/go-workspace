package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/mizanmahi/types"
)

func main() {
	log.Println("Worker started — processing events")
	// In a real app you'd consume from a queue (Kafka, RabbitMQ, etc.)
	// Here we simulate an incoming event every 2 seconds.
	for {
		event := types.Event{
			Type:    types.EventUserCreated,
			Payload: &types.User{ID: "42", Name: "Bob"},
		}
		process(event)
		time.Sleep(2 * time.Second)
	}
}

func process(e types.Event) {
	b, _ := json.MarshalIndent(e, "", "  ")
	log.Printf("Processing event:\n%s\n", b)
}
