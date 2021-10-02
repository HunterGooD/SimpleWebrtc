package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

func CreateHashFromString(s string) string {
	sha := sha256.New()
	return hex.EncodeToString(sha.Sum([]byte(s)))
}

func CreateUUIDString() string {
	return uuid.New().String()
}
