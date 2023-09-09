package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("dev")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("配置初始化失败", err)
	}
}

func InitMysql() {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dns := viper.GetString("mysql.dns")
	var err error
	if db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger,
	}); err != nil {
		fmt.Println("数据库连接失败", err)
	}
}

func GetDb() *gorm.DB {
	return db
}
