package bcrypt_adapter

import (
	"github.com/doga10/clean-architecture-go/src/data/protocols/criptography"
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter interface {
	criptography.Hasher
	criptography.HashComparer
}

type crypt struct {}

func NewBcryptAdapter() BcryptAdapter {
	return &crypt{}
}

func (b *crypt) Hash(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(bytes), err
}

func (b *crypt) Compare(plaintext string, digest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(plaintext))
	return err == nil
}