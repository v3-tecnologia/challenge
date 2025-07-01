package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGyroscope(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gyroscope, err := NewGyroscope(user.ID, "00:1A:2B:3C:4D:5E", 1, 1.1, -1.1, "2025-06-30 10:00:01")
	assert.Nil(t, err)
	assert.NotNil(t, gyroscope)
	assert.NotEmpty(t, gyroscope.ID)
	assert.NotEmpty(t, gyroscope.User)
	assert.Equal(t, "00:1A:2B:3C:4D:5E", gyroscope.MacAddress)
}

func TestGyroscopeWhenMacIsRequired(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gyroscope, err := NewGyroscope(user.ID, "", 1, 1.1, -1.1, "2025-06-30 10:00:01")
	assert.Equal(t, ErrMacIsRequired, err)
	assert.Nil(t, gyroscope)
}

func TestGyroscopeValidate(t *testing.T) {
	user, _ := NewUser("Clark Kent", "clark@kent.com", "krypto")
	gyroscope, err := NewGyroscope(user.ID, "00:1A:2B:3C:4D:5E", 1, 1.1, -1.1, "2025-06-30 10:00:01")
	assert.Nil(t, err)
	assert.NotNil(t, gyroscope)
	assert.Nil(t, gyroscope.Validate())
}
