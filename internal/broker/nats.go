package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

func ConnectNATS(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Printf("Error connecting to NATS: %v", err)
		return nil, err
	}
	log.Println("Connected to NATS")
	return nc, nil
}
