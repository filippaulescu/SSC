package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	//postgresql://ssc_db_user:DZ5pGuhcBXfUwfqmvHk442roWOycMDi8@dpg-cvctfgrv2p9s73c9pe70-a.frankfurt-postgres.render.com/ssc_db
	//postgresql://@dpg-cvctfgrv2p9s73c9pe70-a.frankfurt-postgres.render.com/ssc_db
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Connected to the database successfully!")
}
