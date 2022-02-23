package mysql

import (
	"bluebell_backend/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
