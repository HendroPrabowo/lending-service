package database

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func InitMysql() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if Mysql != nil {
		log.Info("MYSQL : using existing instance")
		return Mysql
	}

	log.Info("MYSQL : create new instance")

	mysqlUrl := os.Getenv("MYSQL_DATABASE_URL")
	if mysqlUrl == "" {
		log.Info("MYSQL : config from localhost")
		mysqlUrl = MYSQL_URL
	}

	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = db

	log.Info("MYSQL : instance created")
	return Mysql
}
