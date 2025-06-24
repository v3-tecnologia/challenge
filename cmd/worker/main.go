package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dryingcore/v3-challenge/internal/config"
	"github.com/dryingcore/v3-challenge/internal/core/model"
	"github.com/nats-io/nats.go"
)

func main() {
	config.Load()
	nc, err := nats.Connect(config.NATSUrl)
	if err != nil {
		log.Fatalf("❌ There is an error while trying to connect on NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("❌ Error while trying to access JetStream: %v", err)
	}

	sub, err := js.PullSubscribe("telemetry.gps", "worker-gps",
		nats.Bind("TELEMETRY", "worker-gps"),
	)
	if err != nil {
		log.Fatalf("❌ An error was ocurred while trying to PullSubscribe: %v", err)
	}

	log.Println("👂 Worker is awake, processing JetStream messages...")
	for {
		msgs, err := sub.Fetch(10, nats.MaxWait(2*time.Second))
		if err != nil {
			log.Printf("⛔ Error while trying to get messages: %v", err)
			continue
		}

		for _, msg := range msgs {
			var data model.GPSData
			if err := json.Unmarshal(msg.Data, &data); err != nil {
				log.Printf("❌ Failed to parse GPS: %v", err)
				msg.Nak()
				continue
			}

			log.Printf("📍GPS RECEIVED: %+v", data)
			msg.Ack()
		}
	}
}
