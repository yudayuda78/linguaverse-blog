package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Ambil variabel dari env
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    sslMode := os.Getenv("DB_SSLMODE")

    // Buat connection string
    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, dbName, sslMode,
    )

    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        panic(err)
    }

    // Test koneksi
    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Database connected successfully!")
}