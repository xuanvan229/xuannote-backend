package config

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-ini/ini"
	"log"
)

// App config
type App struct {
	PageSize int
	JwtSecret string
	PrefixURL string
	RuntimeRootPath string
	ImageSavePath string
	ImageAllowExts []string
}

// Server config
type Server struct {
	HTTPPort int
}

// Postgres config
type Postgres struct {
	User string
	Password string
	Host string
	Port string
	DatabaseName string
}

// Redis config
type Redis struct {
	Host string
	Password string
	NetWork string
	Size int
	KeyPairs string
}

var (
	//AppSetting config
	AppSetting = &App{}

	// ServerSetting config
	ServerSetting = &Server{}

	// PostgresSetting config
	PostgresSetting = &Postgres{}

	// RedisSetting config
	RedisSetting = &Redis{}

 	cfg *ini.File
)


var (
	host     = "localhost"
	port     = "5434"
	user     = "postgres"
	password = "k8kwQ8f4A2fjZk3QhyebekRYKK"
	dbname   = "flicker"

	PrefixUrl  = "http://localhost:1323"
	RuntimeRootPath = "runtime"
	ImageSavePath = "upload/images"
)


func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

// Setup app, loading app.ini to create setup project.
func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("postgres", PostgresSetting)
	mapTo("redis", RedisSetting)
}

// Connect to database
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host="+PostgresSetting.Host+" port="+PostgresSetting.Port+" user="+PostgresSetting.User+" dbname="+PostgresSetting.DatabaseName+" password="+PostgresSetting.Password+" sslmode=disable")
	if err != nil {
		return db, err
	}
	return db, nil
}
