package bb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES256GCMEncrypt(t *testing.T) {
	m1 := []byte("plaintext")
	k := KDF([]byte("password"), nil)

	r := bytes.NewReader(m1)
	w := new(bytes.Buffer)

	err := AES256GCMEncrypt(r, w, k)
	assert.Nil(t, err)

	m2 := w.Bytes()
	assert.NotNil(t, m2)

	r = bytes.NewReader(m2)
	w = new(bytes.Buffer)

	err = AES256GCMDecrypt(r, w, k)
	assert.Nil(t, err)
}

func TestAES256GCMEncryptBytes(t *testing.T) {
	m1 := []byte("plaintext")

	k := KDF([]byte("password"), nil)
	c, err := SymmetricEncryptBytes(m1, k, AES256GCM)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	m2, err := SymmetricDecryptBytes(c, k, AES256GCM)
	assert.Nil(t, err)
	assert.NotNil(t, m2)

	assert.Equal(t, m1, m2)
}
