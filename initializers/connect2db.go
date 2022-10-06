package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func Connect2DB() {
	dsn := "host=localhost dbname=grpcdb port=5432 sslmode=disable"
	config := &gorm.Config{}
	//config.Logger = logger.Default.LogMode(logger.Info)

	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
