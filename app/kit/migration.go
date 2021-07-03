package kit

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

func loadMigrationsDir() *migrate.FileMigrationSource {
	var dir = os.Getenv("migrations")
	if dir == "" {
		panic(".env migrations is empty")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: path.Join(Env.Dir, dir),
	}
	return migrations
}

func RunMigration(db *sql.DB, direciton migrate.MigrationDirection) {
	dir := loadMigrationsDir()

	n, err := migrate.ExecMax(db, "postgres", dir, direciton, 9999)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func Migration(db *sql.DB) {
	dir := loadMigrationsDir()

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

	n, err := migrate.ExecMax(db, "postgres", dir, direction, space)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	if onlyMigrate {
		fmt.Println("Only run migrate, Done!")
		os.Exit(0)
	}
}
