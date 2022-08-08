package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() *gorm.DB {
	var db *gorm.DB
	var err error

	if e := godotenv.Load(); e != nil {
		log.Println(e)
	}

	//database info
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbName := os.Getenv("PSQL_DB")
	host := os.Getenv("PSQL_HOST")
	sslMode := "require"
	if os.Getenv("DEVELOPMENT") == "true" {
		sslMode = "disable"
	}

	//database uri
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbName, sslMode, password)

	//connect to database
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(
			postgres.Open(dbUri),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
		)
		if err != nil {
			log.Println("Connected to database host.")
			break
		}
		log.Println(err, "\nReconnecting....")
		time.Sleep(3 * time.Second)
	}

	return db
}
