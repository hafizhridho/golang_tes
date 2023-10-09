package configs

import (
	"berita/models/news"
	"fmt"
	"os"
	

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Deklarasikan variabel DB di tingkat paket

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
}
}
func InitDatabase() {
	host := os.Getenv("PGHOST")
    user := os.Getenv("PGUSER")
    password := os.Getenv("PGPASSWORD")
    database := os.Getenv("PGDATABASE")
    port := os.Getenv("PGPORT")
    sslmode := os.Getenv("PGSSLMODE")
    timezone := os.Getenv("PGTIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	host, user, password, database, port, sslmode, timezone)

	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("Failed connect to Database")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&news.News{})
}
