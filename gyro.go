package main

import (
	"github.com/nats-io/nats.go"
	"io"
	"log"
	"net/http"
)

type GyroData struct {
	UniqueId  string  `json:"unique_id"`
	Timestamp string  `json:"timestamp"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

func (ctx *AppContext) buildGyroSchema() {
	var err error
	ctx.GyroSchema, err = ctx.Compiler.Compile([]byte(`{
		"type": "object",
		"properties": {
			"unique_id": { "type": "string", "minLength": 1 },
			"timestamp": { "type": "string", "minLength": 1 },
			"x": { "type": "number" },
			"y": { "type": "number" },
			"z": { "type": "number" }
		},
		"required": [ "unique_id", "timestamp", "x", "y", "z" ]
	}`))
	if err != nil {
		log.Print("could not build GyroSchema")
		log.Fatal(err)
	}
}

// @Summary		envia dados de giroscópio
// @Description	valida e persiste dados de giroscópio
// @Accept			json
// @Produce		json
// @Param			gps	body		GyroData	true	"dados do giroscópio"
// @Success		200	{object}	map[string]string
// @Failure		400	{string}	string	"erro de validação"
// @Failure		405	{string}	string	"método não permitido"
// @Router			/telemetry/gyroscope [post]
func (ctx *AppContext) gyroHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("gyro endpoint called")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("could not read request body")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Printf("gyro body: %s", string(bodyBytes))

	result := ctx.GyroSchema.Validate(bodyBytes)
	if result.IsValid() {
		err = ctx.NC.Publish("gyro", bodyBytes)
		if err != nil {
			http.Error(w, "Failed to Publish", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	} else {
		log.Print("gyro schema validation failed")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (ctx *AppContext) gyroConsumer() error {
	_, err := ctx.NC.Subscribe("gyro", func(msg *nats.Msg) {
		var gyroData GyroData
		err := ctx.GyroSchema.Unmarshal(&gyroData, msg.Data)
		if err != nil {
			log.Print(err)
			return
		}

		log.Println(gyroData)

		_, err = ctx.insertGyro(gyroData)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print("gyro data saved")
	})
	return err
}
