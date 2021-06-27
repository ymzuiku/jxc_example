package kit

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// 确保位数一样的字符串，适合用于短信验证码
func RandomCode(digits int) string {
	min := int(math.Pow(10, float64(digits))) - 1
	max := min*10 + 9
	fmt.Println(min, max)
	if Env.IsDev {
		return "999999"
	}
	num := rand.Intn(max) + min
	label := strconv.Itoa(num)
	out := label[0 : len(label)-1]
	return out
}
