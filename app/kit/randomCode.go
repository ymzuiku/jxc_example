package kit

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 确保位数一样的字符串，适合用于短信验证码
func RandomCode(width int) string {
	if Env.IsDev {
		return "999999"
	}
	return RandomCodeBase(width)
}

func RandomCodeBase(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()

}
