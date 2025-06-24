package main

import (
	"log"

	"github.com/dryingcore/v3-challenge/internal/config"
	"github.com/nats-io/nats.go"
)

func main() {
	config.Load()
	nc, err := nats.Connect(config.NATSUrl)
	if err != nil {
		log.Fatalf("❌ Error while connecting to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("❌ Error while tried to access JetStream: %v", err)
	}

	streamName := "TELEMETRY"
	consumerName := "worker-gps"

	if _, err := js.StreamInfo(streamName); err != nil {
		log.Printf("⚠️ Stream '%s' not found, creating...", streamName)
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{"telemetry.*"},
			Storage:  nats.FileStorage,
		})
		if err != nil {
			log.Fatalf("❌ Error while tried to create stream: %v", err)
		}
		log.Println("✅ Stream created.")
	} else {
		log.Println("✅ Stream already existis, skipping creation.")
	}

	if _, err := js.ConsumerInfo(streamName, consumerName); err != nil {
		log.Printf("⚠️ Consumer '%s' not found, creating...", consumerName)
		_, err := js.AddConsumer(streamName, &nats.ConsumerConfig{
			Durable:       consumerName,
			AckPolicy:     nats.AckExplicitPolicy,
			FilterSubject: "telemetry.gps",
		})
		if err != nil {
			log.Fatalf("❌ Error while tried to create consumer: %v", err)
		}
		log.Println("✅ Consumer created.")
	} else {
		log.Println("✅ Consumer already exists, skipping creation.")
	}
}
