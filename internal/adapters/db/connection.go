package mysql

import (
	"time"

	"github.com/ThePratikSah/flag-zero/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection(cfg *config.Config) *Connection {
	db, err := gorm.Open(mysql.Open(cfg.Database.MySQLDSN), &gorm.Config{})
	config.Check(err)

	sqlDB, err := db.DB()
	config.Check(err)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return &Connection{DB: db}
}

func (c *Connection) Close() error {
	sqlDB, err := c.DB.DB()
	config.Check(err)
	return sqlDB.Close()
}

func (c *Connection) GetDB() *gorm.DB {
	return c.DB
}
