package config

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var (
	Cfg *ini.File
	dbType string
	dbUser string
	RunMode string
	dbPassword string
	dbHost string
	dbName string
	Port int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init()  {
	var err error
	Cfg, err = ini.Load("config/app.ini")
		if err != nil {
			log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	ConnectDb()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func ConnectDb()  {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").MustString("mysql")
	dbUser = sec.Key("USER").MustString("jimmy")
	dbPassword = sec.Key("PASSWORD").MustString("")
	dbHost = sec.Key("HOST").MustString("")
	dbName = sec.Key("NAME").MustString("shop")

	connect, err := sqlx.Open(dbType, dbUser + dbPassword + "@tcp(" + dbHost + ")/" + dbName)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println(connect)
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Port = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	ReadTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}