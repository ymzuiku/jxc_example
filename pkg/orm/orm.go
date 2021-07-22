package orm

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ymzuiku/gewu_jxc/pkg/env"
	"gorm.io/driver/postgres"
)

var SqlDB *sql.DB

var DB *gorm.DB

func gormLog() logger.Interface {
	isDev := env.IsDev

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

	if env.IgnoreSQLLog {
		logConf = logger.Config{
			SlowThreshold:             time.Millisecond * 200, // 慢 SQL 阈值
			LogLevel:                  logger.Error,           // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略记录器的 ErrRecordNotFound 错误
			Colorful:                  true,                   // 颜色
		}
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logConf,
	)
	return newLogger
}

func Init() {
	var err error

	if DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_CONNECT_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		// PrepareStmt:     true,
		CreateBatchSize: 1000,
		Logger:          gormLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		log.Fatalln(err)
	}

	if SqlDB, err = DB.DB(); err != nil {
		log.Fatalln(err)
	}

	maxOpenConns, _ := strconv.Atoi(os.Getenv("maxOpenConns"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("maxIdleConns"))
	maxLifetime, _ := strconv.Atoi(os.Getenv("maxLifetime"))
	SqlDB.SetMaxOpenConns(maxOpenConns)
	SqlDB.SetMaxIdleConns(maxIdleConns)
	SqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
}

var ErrRowsAffected = errors.New("Db RowsAffected is zero")

// 若 gorm.DB 有错误，或 RowsAffected == 0, 都会返回错误
func Ok(tx *gorm.DB) error {
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrRowsAffected
	}
	return nil
}
