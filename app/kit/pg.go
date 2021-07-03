package kit

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
)

var Db *sql.DB
var ORM *gorm.DB

func gormLog() logger.Interface {
	isDev := Env.IsDev

	logConf := logger.Config{
		SlowThreshold:             time.Second,   // 慢 SQL 阈值
		LogLevel:                  logger.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略记录器的 ErrRecordNotFound 错误
		Colorful:                  false,         // 颜色
	}

	if isDev {
		logConf = logger.Config{
			SlowThreshold:             time.Millisecond * 200, // 慢 SQL 阈值
			LogLevel:                  logger.Info,            // 日志级别
			IgnoreRecordNotFoundError: false,                  // 忽略记录器的 ErrRecordNotFound 错误
			Colorful:                  true,                   // 颜色
		}
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logConf,
	)
	return newLogger
}

func PgInit() {
	var err error

	ORM, err = gorm.Open(postgres.Open(os.Getenv("DB_CONNECT_URL")), &gorm.Config{
		CreateBatchSize: 65535,
		Logger:          gormLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}
	Db, err = ORM.DB()

	if err != nil {
		log.Fatalln(err)
	}

	maxOpenConns, _ := strconv.Atoi(os.Getenv("maxOpenConns"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("maxIdleConns"))
	maxLifetime, _ := strconv.Atoi(os.Getenv("maxLifetime"))
	Db.SetMaxOpenConns(maxOpenConns)
	Db.SetMaxIdleConns(maxIdleConns)
	Db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
}
