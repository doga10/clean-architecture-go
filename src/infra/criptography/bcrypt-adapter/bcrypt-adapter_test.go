package bcrypt_adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBcryptAdapter(t *testing.T) {
	crypt := NewBcryptAdapter()
	assert.NotNil(t, crypt)
}

func TestCrypt_Hash(t *testing.T) {
	crypt := NewBcryptAdapter()
	assert.NotNil(t, crypt)
	result, err := crypt.Hash("hash")
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestCrypt_Compare(t *testing.T) {
	crypt := NewBcryptAdapter()
	assert.NotNil(t, crypt)
	hash := "password"
	result, err := crypt.Hash(hash)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	isValid := crypt.Compare(hash, result)
	assert.NotNil(t, isValid)
	assert.True(t, isValid)
}