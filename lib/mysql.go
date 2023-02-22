package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Db           *gorm.DB
	MysqlConnStr = "root:xiet,12345@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	d, err := gorm.Open(mysql.Open(MysqlConnStr), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	Db = d
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
