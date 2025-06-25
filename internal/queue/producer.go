package queue

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var conn *nats.Conn

func InitNATS() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL // nats://127.0.0.1:4222 para local
	}
	log.Printf("Tentando conectar no NATS em: %s", natsURL)

	var err error
	for i := 1; i <= 10; i++ {
		conn, err = nats.Connect(natsURL)
		if err == nil {
			log.Println("Conectado ao NATS com sucesso.")
			return
		}
		log.Printf("Tentativa %d: erro ao conectar no NATS: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("Falha ao conectar no NATS após várias tentativas: %v", err)
}

func PublishImage(data []byte) error {
	return conn.Publish("telemetry.image", data)
}
