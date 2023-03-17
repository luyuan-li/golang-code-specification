package pkg

import (
	"fmt"

	"github.com/golang-code-specification/internal/app/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func InitMysqlDB(cfg config.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("start mysql client failed, db:%s, err:%s", cfg.Database, err.Error())
	}

	mysqlDB = db
	return db
}

func GetMysqlDB() *gorm.DB {
	return mysqlDB
}
