package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	log.Info("Building DB configurations...")
	dbConfig := DBConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",
		Password: "password",
		DBName: "figurer",
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
	 	dbConfig.Password,
	 	dbConfig.Host,
	 	dbConfig.Port,
	 	dbConfig.DBName,
	)
}

