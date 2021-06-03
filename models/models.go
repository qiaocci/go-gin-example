package models

import (
	"fmt"
	"github.com/qiaocci/go-gin-example/pkg/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)
	sec, err := settings.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "failed to get section 'database': %v", err)
	}
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbName)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 日志等级
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,        // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB() // 常规数据库接口
	if err != nil {
		log.Printf("获取常规数据库接口异常, err:%v\n", err)
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n\n", err)
		return
	}
	fmt.Println("ping ok")
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	db.AutoMigrate(&Tag{})
}
