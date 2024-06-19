package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Matheus", "a@a.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Matheus", user.Name)
	assert.Equal(t, "a@a.com", user.Email)
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := NewUser("Matheus", "a@a.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("123456"))
	assert.False(t, user.CheckPassword("654321"))
}
