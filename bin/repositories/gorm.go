package repositories

import (
	"fmt"
	"time"

	c "todo/bin/config"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init(dialector gorm.Dialector) *gorm.DB {
	var db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic("Failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func Mysql() gorm.Dialector {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Env.MysqlUsername,
		c.Env.MysqlPassword,
		c.Env.MysqlHost,
		c.Env.MysqlPort,
		c.Env.MysqlDbName)
	return mysql.Open(conn)
}
