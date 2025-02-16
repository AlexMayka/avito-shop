package unit_test

import (
	"avito_shop/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	secret := "test-secret"
	userId := uint(1)
	username := "testuser"

	token, err := pkg.GenerateJWT(userId, username, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	id, user, err := pkg.ValidateJWT(token, secret)
	assert.NoError(t, err)
	assert.Equal(t, userId, id)
	assert.Equal(t, username, user)
}
