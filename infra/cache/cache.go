package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"sync"
)

var (
	memoryCache = make(map[string]bool)
	mu          sync.RWMutex
)

func hashImage(image []byte) string {
	sum := sha256.Sum256(image)
	return hex.EncodeToString(sum[:])
}

func IsCached(image []byte) (bool, bool) {
	key := hashImage(image)

	mu.RLock()
	val, exists := memoryCache[key]
	mu.RUnlock()
	log.Printf("Cache lookup key=%s exists=%v val=%v\n", key, exists, val)
	return val, exists
}

func SetCache(image []byte, recognized bool) {
	key := hashImage(image)

	mu.Lock()
	memoryCache[key] = recognized
	log.Printf("Cache set key=%s recognized=%v\n", key, recognized)
	mu.Unlock()
}
