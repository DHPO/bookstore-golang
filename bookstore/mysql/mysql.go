package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(v *viper.Viper) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", v.GetString("mysql.username"), v.GetString("mysql.password"), v.GetString("mysql.host"), v.GetString("mysql.port"), v.GetString("mysql.db"))
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}
