package config

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() {
	InitViper()
	InitDB()
	//InitRedis()
}

func InitViper() {
	viper.SetConfigFile("./config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %"))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file is changed")
	})
}

func InitDB() {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbName := viper.GetString("database.dbName")
	dbPort := viper.GetString("database.dbPort")
	dbHost := viper.GetString("database.dbHost")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		dbHost,
		dbPort,
		dbName,
	)

	model.Database(dns)
}

func InitRedis() {
	addr := viper.GetString("redis.addr")
	pwd := ""
	db := viper.GetInt("redis.db")

	cache.Redis(addr, pwd, db)
}
