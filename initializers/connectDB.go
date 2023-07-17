package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := "postgres://admin:IdjGaAQQypDg6fniS8ziZLUD0nPnKSnV@dpg-cilqjctph6eg6kamctpg-a/easytickets"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect To The Database")
	}

	installExtension := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""
	err = DB.Exec(installExtension).Error
	if err != nil{
		log.Fatal("Failed to install uuid-ossp extension")
	}

	fmt.Println("? Connected Successfully To The Database")
}