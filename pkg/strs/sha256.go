package strs

import (
	"crypto/sha256"
	"fmt"

	"github.com/ymzuiku/gewu_jxc/pkg/env"
)

func Sha256(text string) string {
	h := sha256.Sum256([]byte(text + env.Sha256Slat))
	out := fmt.Sprintf("%x", h)
	return out
}
