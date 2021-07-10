package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIDs(t *testing.T) {
	t.Run("get ids", func(t *testing.T) {
		type Dog struct {
			Age int32
		}

		var dogs []Dog
		dogs = append(dogs, Dog{Age: 20})
		dogs = append(dogs, Dog{Age: 30})
		ids, err := GetIDs(&dogs, "Age")

		assert.Nil(t, err)
		assert.Equal(t, ids[0], int32(20))
		assert.Equal(t, ids[1], int32(30))
	})
}
