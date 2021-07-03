package kit

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	t.Run("load env", func(t *testing.T) {
		EnvInit()
		if EnvDir == "" {
			t.Error("No have Env Dir")
		}
		if os.Getenv("DB_CONNECT_URL") == "" {
			t.Error("No have DB_CONNECT_URL")
		}
	})
}
