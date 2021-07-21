package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	t.Run("load env", func(t *testing.T) {
		Init()
		assert.True(t, IsDev)
		assert.NotEmpty(t, RootDir)
		assert.NotEmpty(t, os.Getenv("DB_CONNECT_URL"))
	})
}
