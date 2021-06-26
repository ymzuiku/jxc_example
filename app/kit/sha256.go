package kit

import (
	"crypto/sha256"
	"fmt"
)

var salt = "gewu_sha_salt_jxc"

func Sha256(text string) string {
	h := sha256.Sum256([]byte(text + salt))
	out := fmt.Sprintf("%x", h)
	return out
}
