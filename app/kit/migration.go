package kit

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

func Migration(db *sql.DB) {
	var dir = os.Getenv("migrations")
	if dir == "" {
		panic(".env migrations is empty")
	}
	var isNeedMigrate = false
	var direction migrate.MigrationDirection
	var space int
	var onlyMigrate = false

	if os.Getenv("onlyMigrate") != "" {
		onlyMigrate = true
	}
	var fixConfig = func(key string) {
		if isNeedMigrate {
			return
		}

		value := os.Getenv(key)

		if value == "" {
			return
		}

		isNeedMigrate = true

		if key == "upMigrate" {
			direction = migrate.Up
		} else {
			direction = migrate.Down
		}

		var err error
		space, err = strconv.Atoi(value)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fixConfig("upMigrate")
	fixConfig("downMigrate")

	if !isNeedMigrate {
		fmt.Println("No need migrate.")
		return
	}

	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.ExecMax(db, "postgres", migrations, direction, space)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	if onlyMigrate {
		fmt.Println("Only run migrate, Done!")
		os.Exit(0)
	}
}
