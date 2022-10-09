package helpers

import (
	"database/sql"
	"final-project/configs"
	"final-project/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn string
	DB   *gorm.DB
	err  error
)

func PgSqlDB(envConfig configs.Database) *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", envConfig.Host, envConfig.User, envConfig.Password, envConfig.Port, envConfig.Name)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		conf := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", envConfig.Host, envConfig.User, envConfig.Password, envConfig.Port)
		db, err := sql.Open("postgres", conf)

		if err != nil {
			log.Fatal(err)
		} else {
			_, err = db.Exec("create database " + envConfig.Name)
			if err != nil {
				log.Fatal(err)
			} else {
				DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
			}
		}
	}

	DB.Debug().AutoMigrate(&models.User{})
	return DB
}
