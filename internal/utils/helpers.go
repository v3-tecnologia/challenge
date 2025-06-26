package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strings"
	"time"
)

// IsValidImageBase64 verifica se uma string é um Base64 válido
func IsValidImageBase64(s string) bool {
	if s == "" {
		return false
	}

	// Remover data URL prefix se presente (ex: "data:image/png;base64,")
	if strings.HasPrefix(s, "data:image/") {
		parts := strings.Split(s, ",")
		if len(parts) != 2 {
			return false
		}
		s = parts[1]
	}

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}

	// Verifica se os dados decodificados formam uma imagem válida
	_, _, err = image.Decode(bytes.NewReader(decoded))
	return err == nil
}

// ParseTimestamp converte string de timestamp para time.Time
func ParseTimestamp(timestampStr string) (time.Time, error) {
	// Tenta diferentes formatos de timestamp
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999Z07:00",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.999999",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timestampStr); err == nil {
			return t, nil
		}
	}

	// Se nenhum formato funcionou, retorna erro
	return time.Time{}, ErrInvalidTimestampFormat
}

func GetZeroTime() time.Time {
	return time.Time{}
}
