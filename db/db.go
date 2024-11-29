package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("POSTGRES_DNS")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	fmt.Println("Connected to the database successfully!")
}
