package config

import (
	"log"
	"time"

	"github.com/lv123123long/pine/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initDB 初始化数据库连接
func initDB() {
	// 数据库连接字符串
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取底层的 sql.DB 实例
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB instance: %v", err)
	}

	// 设置数据库连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	// 如果设置连接池参数时发生错误，记录日志并退出程序
	if err != nil {
		log.Fatalf("Failed to set connection pool parameters: %v", err)
	}

	global.Db = db
}
