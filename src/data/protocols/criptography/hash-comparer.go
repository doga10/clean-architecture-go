package criptography

type HashComparer interface {
	Compare(plaintext string, digest string) bool
}
