package tools

import (
	"math/rand"
	"strconv"
)

// 确保位数一样的字符串，适合用于短信验证码
func RandomCode(max int) string {
	theMax := max*10 + 9
	num := rand.Intn(theMax) + theMax
	label := strconv.Itoa(num)
	out := label[1 : len(label)-1]
	return out
}
