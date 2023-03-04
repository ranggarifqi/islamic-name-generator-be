package main

import (
	"github.com/ranggarifqi/islamic-name-generator-be/helper"
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
)

func main() {
	helper.InitializeEnv("../.env")

	_, _, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

}
