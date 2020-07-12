package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect("mongodb://127.0.0.1:27017", "app")
	assert.Nil(t, err)
	assert.NotNil(t, Db)
	assert.NotNil(t, Client)
	assert.NotNil(t, DbName)
	assert.NotNil(t, Url)
}

func TestConnectError(t *testing.T) {
	err := Connect("", "")
	assert.NotNil(t, err)
}

func TestGetCollection(t *testing.T) {
	err := Connect("mongodb://127.0.0.1:27017", "app")
	assert.Nil(t, err)

	collection := GetCollection("accounts")
	assert.NotNil(t, collection)
}
