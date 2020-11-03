package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(dns string) {
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Printf("database open failed, err: %v", err)
	}

	db.LogMode(true)

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(time.Hour)

	DB = db
	DB.AutoMigrate(&User{}).AutoMigrate(&Video{})
}
