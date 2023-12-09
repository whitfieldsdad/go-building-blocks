package bb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	publicKey, privateKey, err := GenerateRSAKeyPair(2048)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, publicKey, "Public key should not be nil")
	assert.NotNil(t, privateKey, "Private key should not be nil")
}

func TestRSAEncrypt(t *testing.T) {
	publicKey, privateKey, err := GenerateRSAKeyPair(2048)
	assert.Nil(t, err, "Error should be nil")

	m1 := []byte("Hello, World!")
	r := bytes.NewReader(m1)
	w := new(bytes.Buffer)

	// Encrypt
	err = RSAEncrypt(r, w, publicKey)
	assert.Nil(t, err, "Error should be nil")

	// Decrypt
	r = bytes.NewReader(w.Bytes())
	w = new(bytes.Buffer)
	err = RSADecrypt(r, w, privateKey)
	assert.Nil(t, err, "Error should be nil")

	m2 := w.Bytes()
	assert.Equal(t, m1, m2, "Decrypted bytes should match original bytes")
}

func TestRSAEncryptBytes(t *testing.T) {
	publicKey, privateKey, err := GenerateRSAKeyPair(2048)
	assert.Nil(t, err, "Error should be nil")

	m1 := []byte("Hello, World!")
	c, err := RSAEncryptBytes(m1, publicKey)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, c, "Encrypted bytes should not be nil")

	m2, err := RSADecryptBytes(c, privateKey)
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, m2, "Decrypted bytes should not be nil")
	assert.Equal(t, m1, m2, "Decrypted bytes should match original bytes")
}
