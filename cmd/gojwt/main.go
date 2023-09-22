package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/monirz/gojwt/config"
	"github.com/monirz/gojwt/handler"

	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//get the env variables
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
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(dsn)
		panic(err)
	}

	s := handler.NewServer(db)
	s.Config = cfg

	s.Run()
}
