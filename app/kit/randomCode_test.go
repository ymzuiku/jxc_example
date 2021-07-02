package kit

import (
	"strconv"
	"testing"
)

func TestRandomCode(t *testing.T) {

	t.Run("len", func(t *testing.T) {
		Env.IsDev = false

		a := RandomCode(5)
		if len(a) != 5 {
			t.Errorf("RandomCode 不等于5")
		}

		a = RandomCode(1)
		if len(a) != 1 {
			t.Errorf("RandomCode 不等于1")
		}

		a = RandomCode(10)
		if len(a) != 10 {
			t.Errorf("RandomCode 不等于10")
		}

	})

	t.Run("is like int", func(t *testing.T) {
		a := RandomCode(6)
		_, err := strconv.Atoi(a)
		if err != nil {
			t.Errorf("RandomCode 不是纯数字")
		}

		a = RandomCode(1)
		_, err = strconv.Atoi(a)
		if err != nil {
			t.Errorf("RandomCode 不是纯数字 长度0")
		}

	})

	t.Run("randomCode test 999999", func(t *testing.T) {
		Env.IsDev = true

		a := RandomCode(10)
		if a != "999999" {
			t.Errorf("RandomCode 测试环境不等于999999")
		}

		a = RandomCode(6)
		if a != "999999" {
			t.Errorf("RandomCode 测试环境不等于999999")
		}

		a = RandomCode(10)
		i, err2 := strconv.Atoi(a)
		if err2 != nil {
			t.Errorf("RandomCode 测试环境转数字错误")
		}
		if i != 999999 {
			t.Errorf("RandomCode 测试环境不等于999999")
		}

	})

}
