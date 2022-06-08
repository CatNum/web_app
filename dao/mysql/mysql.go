package mysql

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web_app/settings"
)

var DB *gorm.DB
var sqlDB *sql.DB

func Init(cfg *settings.MySQLConf)(err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		zap.L().Error("connect DB failed,err:%v\n",zap.Error(err))
		return err
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err = DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 初始化用户表
	err = InitUser()
	// 初始化社区表
	err = InitCommunity()
	//初始化社区详情表
	err = InitCommunityDetail()
	return err
}

func Close(){
	_ = sqlDB.Close()
}
