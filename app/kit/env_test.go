package kit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	t.Run("load env", func(t *testing.T) {
		EnvInit()
		assert.True(t, Env.IsDev)
		assert.NotEmpty(t, Env.RootDir)
		assert.NotEmpty(t, os.Getenv("DB_CONNECT_URL"))
	})
}
