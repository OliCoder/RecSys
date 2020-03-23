package models

import (
	"fmt"
	"github.com/OliCoder/RecSys/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	config := settings.GetConfigInstance()
	dbType := config.DBType
	host := config.DBHost
	port := config.DBPort
	dbName := config.DBName
	user := config.User
	password := config.Password
	tablePrefix := config.TablePrefix

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName))
	if err != nil {
		log.Errorf("Fail to init db, err:%s", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}
