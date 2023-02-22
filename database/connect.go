package database

import (
	// "os"
	// "fmt"
	"github.com/vanessamae23/cvwo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

var DB *gorm.DB

func Connect() {

	// url := os.Getenv("DB_URL")
	// db := os.Getenv("DB_DATABASE")
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")

	dsn := "root:makmae@/cvwo"
	
	//dsn := fmt.Sprintf("%s:%s@%s/%s", username, password, url, db)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database")
	}

	DB = database
	// To create tables
	database.AutoMigrate(&models.User{}, &models.Forum{}, &models.Comment{})

}
