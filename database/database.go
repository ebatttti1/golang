package database

import(
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	DB, err := sql.Open("users", dsn)
	if err != nil {
		log.Fatal("error in connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("error in connecting to database:", err)
	}

	fmt.Println("Connected to database")
}








