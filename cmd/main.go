package main

import (
	"github.com/joho/godotenv"
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	_, _, err = mongodb.Connect()
	if err != nil {
		panic(err)
	}

}
