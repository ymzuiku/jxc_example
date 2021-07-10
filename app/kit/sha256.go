package kit

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(text string) string {
	h := sha256.Sum256([]byte(text + Env.Sha256Slat))
	out := fmt.Sprintf("%x", h)
	return out
}
