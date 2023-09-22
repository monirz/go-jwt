package main

import (
	"log"

	"github.com/monirz/gojwt/handler"

	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//get the env variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	s := handler.NewServer()

	s.Run()
}
