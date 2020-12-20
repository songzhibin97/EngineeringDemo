package pkg

import (
	"fmt"

	"gorm.io/driver/mysql"

	"github.com/jmoiron/sqlx"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 各种初始化 cmd 命令

// DataInit: 数据库初始化
func DataInit() (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("DB.User"),
		viper.GetString("DB.Password"),
		viper.GetString("DB.Host"),
		viper.GetInt("DB.Port"),
		viper.GetString("DB.DBName"))
	dbInstantiate, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		dbInstantiate.Close()
		return nil, nil, err
	}
	//db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&biz.User{})
	return db, func() {
		dbInstantiate.Close()
	}, nil
}
