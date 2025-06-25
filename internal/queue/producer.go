package queue

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var conn *nats.Conn

func InitNATS() {
	credsFile := "./NGS-Default-CLI.creds"
	natsURL := os.Getenv("NATS_URL")

	var err error

	if _, errFile := os.Stat(credsFile); errFile == nil {
		if natsURL == "" {
			natsURL = "tls://connect.ngs.global:4222"
		}
		log.Printf("Tentando conectar no NATS em %s com credenciais", natsURL)
		for i := 1; i <= 10; i++ {
			conn, err = nats.Connect(natsURL, nats.UserCredentials(credsFile))
			if err == nil {
				log.Println("Conectado ao NATS com credenciais com sucesso.")
				return
			}
			log.Printf("Tentativa %d: erro ao conectar no NATS com credenciais: %v", i, err)
			time.Sleep(2 * time.Second)
		}
	} else {
		log.Println("Arquivo de credenciais não encontrado, tentando conexão local sem credenciais.")
	}

	if natsURL == "" {
		natsURL = nats.DefaultURL // nats://127.0.0.1:4222
	}
	for i := 1; i <= 10; i++ {
		conn, err = nats.Connect(natsURL)
		if err == nil {
			log.Println("Conectado ao NATS local com sucesso.")
			return
		}
		log.Printf("Tentativa %d: erro ao conectar no NATS local: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("Falha ao conectar no NATS após várias tentativas: %v", err)
}

func PublishImage(data []byte) error {
	return conn.Publish("telemetry.image", data)
}
