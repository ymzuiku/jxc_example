package strs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymzuiku/gewu_jxc/pkg/env"
)

func init() {
	env.IsDev = true
}

func TestRandomCode(t *testing.T) {

	t.Run("len", func(t *testing.T) {

		assert.Equal(t, len(RandomCodeBase(5)), 5)

		assert.Equal(t, len(RandomCodeBase(4)), 4)

		assert.Equal(t, len(RandomCodeBase(10)), 10)

	})

	t.Run("is like int", func(t *testing.T) {
		a := RandomCodeBase(6)
		_, err := strconv.Atoi(a)
		assert.Nil(t, err)

		a = RandomCodeBase(1)
		_, err = strconv.Atoi(a)
		assert.Nil(t, err)

	})

	t.Run("randomCode test 999999", func(t *testing.T) {
		assert.Equal(t, RandomCode(10), "999999")

		assert.Equal(t, RandomCode(6), "999999")

		i, err2 := strconv.Atoi(RandomCode(10))
		assert.Nil(t, err2)
		assert.Equal(t, i, 999999)

	})

}
