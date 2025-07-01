package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGps(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gps, err := NewGps(user.ID, "00:1A:2B:3C:4D:5E", -23.853345114882476, -45.28845610372027, "2025-06-30 10:00:01")
	assert.Nil(t, err)
	assert.NotNil(t, gps)
	assert.NotEmpty(t, gps.ID)
	assert.NotEmpty(t, gps.User)
	assert.Equal(t, "00:1A:2B:3C:4D:5E", gps.MacAddress)
}

func TestGpsWhenMacIsRequired(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gps, err := NewGps(user.ID, "", -23.853345114882476, -45.28845610372027, "2025-06-30 10:00:01")
	assert.Equal(t, ErrMacIsRequired, err)
	assert.Nil(t, gps)
}

func TestGpsValidate(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gps, err := NewGps(user.ID, "00:1A:2B:3C:4D:5E", -23.853345114882476, -45.28845610372027, "2025-06-30 10:00:01")
	assert.Nil(t, err)
	assert.NotNil(t, gps)
	assert.Nil(t, gps.Validate())
}
