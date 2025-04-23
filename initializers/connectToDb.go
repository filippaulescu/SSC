package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	// this ones are the old oness
	//postgresql://ssc_db_user:DZ5pGuhcBXfUwfqmvHk442roWOycMDi8@dpg-cvctfgrv2p9s73c9pe70-a.frankfurt-postgres.render.com/ssc_db
	//postgresql://@dpg-cvctfgrv2p9s73c9pe70-a.frankfurt-postgres.render.com/ssc_db
	// new one:postgresql://neondb_owner:npg_wpzkJAD6n5cg@ep-steep-river-a2nml8ej-pooler.eu-central-1.aws.neon.tech/neondb?sslmode=require
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Connected to the database successfully!")
}
