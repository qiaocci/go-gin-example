package models

import (
	"fmt"
	"github.com/qiaocco/go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/qiaocco/go-gin-example/pkg/logging"
)

var db *gorm.DB

func SetUp() {
	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info), // 日志等级
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,                                // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		logging.Info(err)
	}

	sqlDB, err := db.DB() // 常规数据库接口
	if err != nil {
		logging.Printf("获取常规数据库接口异常, err:%v\n", err)
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n\n", err)
		return
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	//db.AutoMigrate(&Article{})
}
