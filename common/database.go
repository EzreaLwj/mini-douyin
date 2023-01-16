package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mini-douyin/config"
)

var DB *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("数据库连接错误：%s", err)
		return
	}

	// 全局赋值
	DB = db

	log.Printf("数据库连接成功")
}
