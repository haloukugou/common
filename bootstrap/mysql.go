package bootstrap

import (
	"dj/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	conf := config.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)
	db, dErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dErr != nil {
		panic(dErr)
	}
	sqlDB, sErr := db.DB()

	if sErr != nil {
		panic(sErr)
	}
	sqlDB.SetConnMaxLifetime(time.Duration(3600) * time.Second)
	return db
}
