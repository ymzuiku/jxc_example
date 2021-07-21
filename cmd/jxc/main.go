package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ymzuiku/env_migrate"
	"github.com/ymzuiku/gewu_jxc/api"
	"github.com/ymzuiku/gewu_jxc/pkg/env"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
	"github.com/ymzuiku/gewu_jxc/pkg/srv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	env.Init()
	orm.Init()
	rds.Init()
	env_migrate.Auto(orm.SqlDB)
	srv.Init()
	api.Init()
	log.Fatal(srv.Fiber.Listen(":3100"))
}
