package inits

import (
	"fmt"
	"gospacex-pengyilong/srv/basic/config"

	"github.com/Lilong-maker/consul"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func MysqlInit() {
	MysqlConfig := consul.Gen.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConfig.User,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
}
func ViperInit() {
	viper.SetConfigFile("C:\\Users\\Lenovo\\Desktop\\gospacex-pengyilong\\config.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&consul.Gen)
	if err != nil {
		return
	}
	fmt.Println("配置加载成功")
}
