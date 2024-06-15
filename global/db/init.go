package db

import (
	"context"
	"fmt"
	mylogger "github/MovieWebsite/global/logger"
	"github/MovieWebsite/global/setting"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	_ "gorm.io/gorm/schema"
)

var (
	_db *gorm.DB
)

func InitDBEngine() {
	if setting.ServerSetting.RunMode == "debug" {
		mylogger.Logger.LogMode(mylogger.LevelDebug)
	}
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		setting.DatabaseSetting.Username,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.DBName,
		setting.DatabaseSetting.Charset,
		setting.DatabaseSetting.ParseTime)
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       s,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}), &gorm.Config{
		Logger: mylogger.Logger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
		panic("数据库连接配置错误")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf(err.Error())
		panic("数据库连接配置错误")
	}
	sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)
	_db = db
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db2 := _db
	return db2.WithContext(ctx)
}
