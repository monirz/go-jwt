package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/monirz/gojwt/config"
	"github.com/monirz/gojwt/utils"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found in the root")
	}

	cfg := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// // Ensure the database connection is successful
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	uid := uuid.NewString()
	pass, err := utils.HashPassword("password")
	if err != nil {
		log.Fatal(err)
	}
	// Insert the admin user into the "users" table
	_, err = db.Exec(
		"INSERT INTO users (uuid, username, email, password, created_at, updated_at) "+
			"VALUES ($1, $2, $3, $4, $5, $6)",
		uid, "admin", "admin@example.com", pass, time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("test user created successfully!")
}
