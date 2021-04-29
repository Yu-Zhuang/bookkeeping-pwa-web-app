package logic

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	ret := hex.EncodeToString(h.Sum(nil))
	return ret
}
