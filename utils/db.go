package utils

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func openDB() {
	mysqlConf := config.GetGlobalConf().DbConfig
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User,
		mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	log.Info("mdb addr:" + connArgs)
	var err error
	db, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to connect database")
		panic("failed to connect database")
	}
	log.Info("connect success")
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("fetch db connection err %s", err.Error())
		panic("fetch db connection err:" + err.Error())
	}
	log.Info("fetch db connection success")
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConn)                                        //设置最大空闲连接
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConn)                                        //设置最大打开的连接
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxIdleTime * int64(time.Second))) //设置空闲时间为(s)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	dbOnce.Do(openDB)
	return db
}
