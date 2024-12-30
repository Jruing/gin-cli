package utils

import (
	"fmt"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/casbin/casbin/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlClient *gorm.DB
	Enforcer    *casbin.Enforcer
)

func init() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("err")
	}

	// Mysql配置文件
	mysqlHost := viper.GetString("mysql.host")
	mysqlPort := viper.GetInt("mysql.port")
	mysqlUser := viper.GetString("mysql.user")
	mysqlDB := viper.GetString("mysql.db")
	mysqlPasswd := viper.GetString("mysql.password")
	mysqldsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysqlUser, mysqlPasswd, mysqlHost, mysqlPort, mysqlDB)
	// 初始化数据库连接
	MysqlClient, err = gorm.Open(mysql.Open(mysqldsn))
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}

	// 初始化casbin
	adapter, err := gormadapter.NewAdapter("mysql", mysqldsn, true)
	if err != nil {
		panic(err)
	}
	Enforcer, _ = casbin.NewEnforcer("config/rbac_model.conf", adapter)

	//其他
}
