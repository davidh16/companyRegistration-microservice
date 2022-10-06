package initializers

import (
	"companyRegistration-microservice/model"
	"log"
)

func init() {
	Connect2DB()
}

func Migrate() {
	err := DB.AutoMigrate(&model.Company{})
	if err != nil {
		log.Fatalf("could not migrate to the database: %v", err)
	}
}
