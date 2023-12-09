package bb

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

type KDFType string

const (
	PBKDF2 KDFType = "pbkdf2"
)

const (
	PBKDF2Rounds     = 100000
	PBKDF2SaltLength = 8
)

func KDF(password, salt []byte) []byte {
	return PBKDF2_HMAC_SHA256(password, salt, PBKDF2Rounds)
}

func PBKDF2_HMAC_SHA256(password, salt []byte, rounds int) []byte {
	return pbkdf2.Key(password, salt, rounds, 32, sha256.New)
}
