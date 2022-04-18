package handler

import (
	"crypto/sha1"
	"encoding/hex"
)

func encrypt(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}
