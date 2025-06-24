package config

import (
	"os"
	"strconv"
	"time"
)

var (
	AllowedSkew time.Duration
	NATSUrl     string
)

func Load() {
	const defaultSkewSeconds = 10
	val := os.Getenv("GYROSCOPE_ALLOWED_SKEW_SECONDS")
	if val == "" {
		AllowedSkew = time.Duration(defaultSkewSeconds) * time.Second
	} else if seconds, err := strconv.Atoi(val); err == nil && seconds >= 0 {
		AllowedSkew = time.Duration(seconds) * time.Second
	} else {
		AllowedSkew = time.Duration(defaultSkewSeconds) * time.Second
	}

	NATSUrl = os.Getenv("NATS_URL")
	if NATSUrl == "" {
		NATSUrl = "nats://localhost:4222"
	}
}
