package kit

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomCode(t *testing.T) {

	t.Run("len", func(t *testing.T) {
		Env.IsDev = false

		a := RandomCode(5)
		assert.Equal(t, len(a), 5)

		a = RandomCode(1)
		assert.Equal(t, len(a), 1)

		a = RandomCode(10)
		assert.Equal(t, len(a), 10)

	})

	t.Run("is like int", func(t *testing.T) {
		a := RandomCode(6)
		_, err := strconv.Atoi(a)
		assert.Nil(t, err)

		a = RandomCode(1)
		_, err = strconv.Atoi(a)
		assert.Nil(t, err)

	})

	t.Run("randomCode test 999999", func(t *testing.T) {
		Env.IsDev = true

		a := RandomCode(10)
		assert.Equal(t, a, "999999")

		a = RandomCode(6)
		assert.Equal(t, a, "999999")

		a = RandomCode(10)
		i, err2 := strconv.Atoi(a)
		assert.Nil(t, err2)
		assert.Equal(t, i, 999999)

	})

}
