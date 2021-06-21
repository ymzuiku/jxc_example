package tools

import (
	"database/sql"
	"gewu_jxc/sql/db"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var Pg *sql.DB
var ORM *db.Queries

func PgInit() {
	conn, err := sql.Open("pgx", os.Getenv("DB_CONNECT_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	maxOpenConns, _ := strconv.Atoi(os.Getenv("maxOpenConns"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("maxIdleConns"))
	maxLifetime, _ := strconv.Atoi(os.Getenv("maxLifetime"))
	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
	Pg = conn
	ORM = db.New(Pg)
}
