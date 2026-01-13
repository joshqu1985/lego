package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Hmac hmac编码.
func Hmac(key, str string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	if _, err := h.Write([]byte(str)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha256 sha256编码.
func Sha256(str string) (string, error) {
	sh := sha256.New()
	if _, err := sh.Write([]byte(str)); err != nil {
		return "", err
	}

	return hex.EncodeToString(sh.Sum(nil)), nil
}
