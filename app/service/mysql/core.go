package mysql

import (
	"exam/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sync"
	"time"
)

var ins *gorm.DB
var once sync.Once

func GetIns() *gorm.DB {
	once.Do(func() {
		ins = NewConnection()
	})
	return ins
}

func NewConnection() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		config.InsMysql().User,
		config.InsMysql().Password,
		config.InsMysql().Host,
		config.InsMysql().Port,
		config.InsMysql().DB,
		config.InsMysql().Charset,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Print(err.Error())
	}
	return conn
}
