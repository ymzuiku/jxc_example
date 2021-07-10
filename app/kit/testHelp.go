package kit

import (
	"math/rand"
	"os"
	"sync"
	"time"
)

var onceTestInit sync.Once

func testInit() {
	rand.Seed(time.Now().UnixNano())
	EnvInit()
	PgInit()
	RedisInit()
	Env.IsDev = true
}

func TestInit() {
	onceTestInit.Do(testInit)
}

func ExitIf(value bool) {
	if !value {
		os.Exit(0)
	}
}
