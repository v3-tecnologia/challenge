package main

import (
	"github.com/nats-io/nats.go"
	"io"
	"log"
	"net/http"
)

type GpsData struct {
	UniqueId  string  `json:"unique_id"`
	Timestamp string  `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (ctx *AppContext) buildGpsSchema() {
	var err error
	ctx.GpsSchema, err = ctx.Compiler.Compile([]byte(`{
		"type": "object",
		"properties": {
			"unique_id": { "type": "string", "minLength": 1 },
			"timestamp": { "type": "string", "minLength": 1 },
			"latitude": { "type": "number" },
			"longitude": { "type": "number" }
		},
		"required": [ "unique_id", "timestamp", "latitude", "longitude" ]
	}`))
	if err != nil {
		log.Print("could not build GpsSchema")
		log.Fatal(err)
	}
}

// @Summary		envia dados de gps
// @Description	valida e persiste dados de gps
// @Accept			json
// @Produce		json
// @Param			gps	body		GpsData	true	"dados do gps"
// @Success		200	{object}	map[string]string
// @Failure		400	{string}	string	"erro de validação"
// @Failure		405	{string}	string	"método não permitido"
// @Router			/telemetry/gps [post]
func (ctx *AppContext) gpsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("gps endpoint called")

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

	log.Printf("gps body: %s", string(bodyBytes))

	result := ctx.GpsSchema.Validate(bodyBytes)
	if result.IsValid() {
		err = ctx.NC.Publish("gps", bodyBytes)
		if err != nil {
			http.Error(w, "Failed to Publish", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	} else {
		log.Print("gps schema validation failed")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (ctx *AppContext) gpsConsumer() error {
	_, err := ctx.NC.Subscribe("gps", func(msg *nats.Msg) {
		var gpsData GpsData
		err := ctx.GpsSchema.Unmarshal(&gpsData, msg.Data)
		if err != nil {
			log.Print(err)
			return
		}

		log.Println(gpsData)

		_, err = ctx.insertGps(gpsData)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print("gps data saved")
	})
	return err
}
