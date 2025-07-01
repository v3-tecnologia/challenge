package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T){
	user, err := NewUser("Clark Kent", "clark@kent.com", "krypto")
	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Clark Kent", user.Name)
	assert.Equal(t, "clark@kent.com", user.Email)
}

func TestUserValidatePassword(t *testing.T){
	user, err := NewUser("Clark Kent", "clark@kent.com", "krypto")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("krypto"))
	assert.False(t, user.ValidatePassword("kryptonita"))
	assert.NotEqual(t, "krypto", user.Password)
}