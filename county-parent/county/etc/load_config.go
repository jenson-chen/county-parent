package etc

import (
	"github.com/jenson/county/county/etc/identity"
	"github.com/jenson/county/county/global"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var vp *viper.Viper

// LoadConfig load all configuration information
func LoadConfig() {
	loadViperConfig() //加载viper配置
	loadDatabase()    //加载DB配置
	loadRedis()       //加载
}

func loadViperConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../") //从项目路径开始
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	vp = viper.GetViper()
}

func loadDatabase() {
	host := vp.GetString("PostgreSQL.host")
	user := vp.GetString("PostgreSQL.user")
	pwd := vp.GetString("PostgreSQL.password")
	dbname := vp.GetString("PostgreSQL.db")
	port := vp.GetInt("PostgreSQL.port")
	sql := identity.PostgreSQL{
		Host:     host,
		User:     user,
		Password: pwd,
		DBName:   dbname,
		Port:     port,
	}
	dsn := sql.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failure,caused by %v\n", err)
	}
	global.DB = db //将该连接设置为一个全局变量
	log.Printf("loading database successfully,%#v", db)
}

func loadRedis() {
	//REDIS中应该存储用户常用的账号 密码 头像等数据

}
