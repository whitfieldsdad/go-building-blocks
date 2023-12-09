package bb

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

const (
	AES256GCM = "aes-256-gcm"
)

const (
	DefaultSymmetricEncryptionAlgorithm = AES256GCM
)

func SymmetricEncrypt(r io.Reader, w io.Writer, k []byte, algo string) error {
	if algo == "" {
		algo = DefaultSymmetricEncryptionAlgorithm
	}
	if algo == AES256GCM {
		return AES256GCMEncrypt(r, w, k)
	} else {
		return errors.New("unsupported algorithm")
	}
}

func SymmetricDecrypt(r io.Reader, w io.Writer, k []byte, algo string) error {
	if algo == "" {
		algo = DefaultSymmetricEncryptionAlgorithm
	}
	if algo == AES256GCM {
		return AES256GCMDecrypt(r, w, k)
	} else {
		return errors.New("unsupported algorithm")
	}
}

func SymmetricEncryptBytes(m, k []byte, algo string) ([]byte, error) {
	if algo == "" {
		algo = DefaultSymmetricEncryptionAlgorithm
	}
	if algo == AES256GCM {
		return AES256GCMEncryptBytes(m, k)
	} else {
		return nil, errors.New("unsupported algorithm")
	}
}

func SymmetricDecryptBytes(c, k []byte, algo string) ([]byte, error) {
	if algo == "" {
		algo = DefaultSymmetricEncryptionAlgorithm
	}
	if algo == AES256GCM {
		return AES256GCMDecryptBytes(c, k)
	} else {
		return nil, errors.New("unsupported algorithm")
	}
}

func AES256GCMEncrypt(r io.Reader, w io.Writer, k []byte) error {
	m, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	c, err := AES256GCMEncryptBytes(m, k)
	if err != nil {
		return err
	}
	_, err = w.Write(c)
	return err
}

func AES256GCMDecrypt(r io.Reader, w io.Writer, k []byte) error {
	c, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	m, err := AES256GCMEncryptBytes(c, k)
	if err != nil {
		return err
	}
	_, err = w.Write(m)
	return err
}

func AES256GCMEncryptBytes(m, k []byte) ([]byte, error) {
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gcm")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "failed to read random bytes")
	}
	c := gcm.Seal(nonce, nonce, m, nil)
	return c, nil
}

func AES256GCMDecryptBytes(c, k []byte) ([]byte, error) {
	aes, err := aes.NewCipher(k)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gcm")
	}
	nonceSize := gcm.NonceSize()
	nonce, c := c[:nonceSize], c[nonceSize:]
	m, err := gcm.Open(nil, []byte(nonce), []byte(c), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt")
	}
	return m, nil
}
