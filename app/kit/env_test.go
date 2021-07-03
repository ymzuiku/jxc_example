package kit

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	t.Run("load env", func(t *testing.T) {
		EnvInit()
		if Env.IsDev == false {
			t.Error("No have Env isDev")
		}
		if Env.Dir == "" {
			t.Error("No have Env Dir")
		}
		if os.Getenv("DB_CONNECT_URL") == "" {
			t.Error("No have DB_CONNECT_URL")
		}
	})
}
