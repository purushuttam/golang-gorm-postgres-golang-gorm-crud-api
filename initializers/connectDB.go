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
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open("postgres://purushuttam:cH8iCTR5nfpgVd9oOE9wlMMYCvcj9EgX@dpg-cidh30enqqlb62jn9jdg-a.singapore-postgres.render.com/restrauntdb"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
