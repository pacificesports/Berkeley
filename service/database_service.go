package service

import (
	"berkeley/config"
	"berkeley/model"
	"berkeley/utils"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

var dbRetries = 0

func InitializeDB() {
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%s;", config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresDatabase, config.PostgresPort)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		if dbRetries < 15 {
			dbRetries++
			utils.SugarLogger.Errorln("failed to connect database, retrying in 5s... ")
			time.Sleep(time.Second * 5)
			InitializeDB()
		} else {
			utils.SugarLogger.Fatalln("failed to connect database after 15 attempts, terminating program...")
		}
	} else {
		utils.SugarLogger.Infoln("Connected to postgres database")
		db.AutoMigrate(&model.School{}, &model.Tag{}, &model.Email{})
		utils.SugarLogger.Infoln("AutoMigration complete")
		DB = db
	}
}
