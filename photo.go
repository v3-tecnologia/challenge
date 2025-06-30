package main

import (
	"github.com/nats-io/nats.go"
	"io"
	"log"
	"net/http"
)

type PhotoData struct {
	UniqueId  string `json:"unique_id"`
	Timestamp string `json:"timestamp"`
	Photo     string `json:"photo"`
}

func (ctx *AppContext) buildPhotoSchema() {
	var err error
	ctx.PhotoSchema, err = ctx.Compiler.Compile([]byte(`{
		"type": "object",
		"properties": {
			"unique_id": { "type": "string", "minLength": 1 },
			"timestamp": { "type": "string", "minLength": 1 },
			"photo": { "type": "string" }
		},
		"required": [ "unique_id", "timestamp", "photo" ]
	}`))
	if err != nil {
		log.Print("failed to build photo schema")
		log.Fatal(err)
	}
}

// @Summary		envia fotos
// @Description	valida e persiste fotos
// @Accept			json
// @Produce		json
// @Param			gps	body		PhotoData	true	"foto"
// @Success		200	{object}	map[string]string
// @Failure		400	{string}	string	"erro de validação"
// @Failure		405	{string}	string	"método não permitido"
// @Router			/telemetry/photo [post]
func (ctx *AppContext) photoHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("photo endpoint called")

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

	log.Printf("photo body: %s", string(bodyBytes))

	result := ctx.PhotoSchema.Validate(bodyBytes)
	if result.IsValid() {
		err = ctx.NC.Publish("photo", bodyBytes)
		if err != nil {
			http.Error(w, "Failed to Publish", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	} else {
		log.Print("photo schema validation failed")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func (ctx *AppContext) photoConsumer() error {
	_, err := ctx.NC.Subscribe("photo", func(msg *nats.Msg) {
		var photoData PhotoData
		err := ctx.PhotoSchema.Unmarshal(&photoData, msg.Data)
		if err != nil {
			log.Print(err)
			return
		}

		log.Println(photoData)

		_, err = ctx.insertPhoto(photoData)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print("photo data saved")
	})
	return err
}
