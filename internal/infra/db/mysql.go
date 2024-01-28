package db

import (
	"fmt"

	"github.com/totoyk/trial-api-golang/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	DBLoc    string
}

func DBConfig() *Database {
	c := new(Database)

	conf, _ := config.LoadEnv()
	c.Host = conf.DBHost
	c.Port = conf.DBPort
	c.UserName = conf.DBUser
	c.Password = conf.DBPass
	c.DBName = conf.DBName
	c.DBLoc = conf.DBLoc

	return c
}

func NewConn() (*gorm.DB, error) {
	c := DBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.UserName, c.Password, c.Host, c.Port, c.DBName)
	dsn += fmt.Sprintf("?charset=utf8mb4&parseTime=true&loc=%s", c.DBLoc)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
