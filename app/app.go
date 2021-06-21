package app

import (
	"gewu_jxc/app/controllers"
	"gewu_jxc/app/tools"
)

func Run() {
	tools.EnvInit()
	tools.PgInit()
	tools.RedisInit()
	tools.Migration(tools.Pg, "sql/migrations")

	controllers.Run(":3100")
}
