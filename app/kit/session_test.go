package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionCreate(t *testing.T) {
	TestInit()
	var session string
	t.Run("create session", func(t *testing.T) {
		var err error
		session, err = SessionCreate("13333333333", "123456")
		assert.Nil(t, err)
		assert.True(t, len(session) > 10)
	})

	t.Run("check session", func(t *testing.T) {
		val, err := SessionGet(session)
		assert.Nil(t, err)
		assert.Equal(t, val, "123456")
	})

	t.Run("has session", func(t *testing.T) {
		err := SessionHas(session)
		assert.Nil(t, err)
	})

	t.Run("remove session", func(t *testing.T) {
		err := SessionRemove(session)
		assert.Nil(t, err)
	})

	t.Run("check session error", func(t *testing.T) {
		val, err := SessionGet(session + "123")
		assert.NotNil(t, err)
		assert.Equal(t, val, "")
	})

	t.Run("has session error", func(t *testing.T) {
		err := SessionHas(session + "123")
		assert.NotNil(t, err)
	})

	t.Run("remove session error", func(t *testing.T) {
		err := SessionRemove(session + "123")
		assert.NotNil(t, err)
	})

}
