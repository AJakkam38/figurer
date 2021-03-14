package main

import (
	"figurer/app"
	"figurer/app/config"
	"figurer/app/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func main()  {
	config.DB, err = gorm.Open(
		"mysql",
		config.DbURL(config.BuildDBConfig()),
	)

	if err != nil {
		log.Fatal("Database connection failed.\nStatus:", err)
	}
	log.Info("Connected to Database successfully !")
	defer config.DB.Close()
	// Migrate the schema
	config.DB.AutoMigrate(&models.Calculations{})

	config.DB.Create(app.Run())
	log.Info("Updated Calculations table !")
}
