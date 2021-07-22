package testh

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/env"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
)

var onceTestInit sync.Once

func UnitTest() {
	onceTestInit.Do(unitTest)
}

func unitTest() {
	rand.Seed(time.Now().UnixNano())
	errox.Debug = true
	env.IgnoreSQLLog = true
	env.IsDev = true
	env.Init()
	orm.Init()
	rds.Init()
	rds.Client.FlushDB(context.Background())
}
