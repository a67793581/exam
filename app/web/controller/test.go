package controller

import (
	"exam/app/config"
	"exam/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func Test(context *gin.Context) {
	name := context.DefaultQuery("name", "jack")
	context.String(200, fmt.Sprintf("hello %s\n", name))
}

func TestMysql(context *gin.Context) {
	dsn := config.MysqlUser + ":" + config.MysqlPassword + "@tcp(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDB + "?charset=" + config.MysqlCharset + "&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(err)
	}
	s := model.Student{Name: "carlo", Key: "test"}
	db.Create(&s)
	db.Delete(&s)
}
