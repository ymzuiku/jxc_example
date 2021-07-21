package main

import (
	"math/rand"
	"time"

	"github.com/ymzuiku/gewu_jxc/pkg/env"
	"github.com/ymzuiku/gewu_jxc/pkg/testh"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	env.IgnoreSQLLog = true
	testh.UnitTest()
	Seeds()
}
