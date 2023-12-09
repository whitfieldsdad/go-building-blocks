package bb

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io"

	"github.com/pkg/errors"
)

const (
	RSA             = "rsa"
	RSA_OAEP_SHA512 = "rsa-oaep-sha512"
)

const (
	DefaultAsymmetricEncryptionAlgorithm = RSA
)

func AsymmetricEncryptBytes(m, k []byte, algo string) ([]byte, error) {
	if algo == "" {
		algo = DefaultAsymmetricEncryptionAlgorithm
	}
	if algo == RSA || algo == RSA_OAEP_SHA512 {
		return RSAEncryptBytes(m, k)
	} else {
		return nil, errors.New("unsupported algorithm")
	}
}

func AsymmetricDecryptBytes(c, k []byte, algo string) ([]byte, error) {
	if algo == "" {
		algo = DefaultAsymmetricEncryptionAlgorithm
	}
	if algo == RSA || algo == RSA_OAEP_SHA512 {
		return RSADecryptBytes(c, k)
	} else {
		return nil, errors.New("unsupported algorithm")
	}
}

func RSAEncrypt(r io.Reader, w io.Writer, publicKey []byte) error {
	m, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	c, err := RSAEncryptBytes(m, publicKey)
	if err != nil {
		return err
	}
	_, err = w.Write(c)
	return err
}

func RSADecrypt(r io.Reader, w io.Writer, privateKey []byte) error {
	c, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	m, err := RSADecryptBytes(c, privateKey)
	if err != nil {
		return err
	}
	_, err = w.Write(m)
	return err
}

func RSAEncryptBytes(m, publicKey []byte) ([]byte, error) {
	k, err := DecodeRSAPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha512.New(), rand.Reader, k, m, nil)
}

func RSADecryptBytes(c, privateKey []byte) ([]byte, error) {
	k, err := DecodeRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	plaintext, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, k, c, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func GenerateRSAKeyPair(bits int) ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyPEMBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPEM,
	}
	return pem.EncodeToMemory(publicKeyPEMBlock), pem.EncodeToMemory(privateKeyPEM), nil
}

func DecodePEM(key []byte) (*pem.Block, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}
	return block, nil
}

func DecodeRSAPublicKey(key []byte) (*rsa.PublicKey, error) {
	block, err := DecodePEM(key)
	if err != nil {
		return nil, err
	}
	derDecodedKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return derDecodedKey, nil
	}
	pemDecodedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := pemDecodedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to decode RSA public key")
	}
	return rsaKey, nil
}

func DecodeRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, err := DecodePEM(key)
	if err != nil {
		return nil, err
	}
	derDecodedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return derDecodedKey, nil
	}
	pemDecodedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := pemDecodedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("failed to decode RSA private key")
	}
	return rsaKey, nil
}
