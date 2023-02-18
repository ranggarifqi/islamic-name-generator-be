package main

import (
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
)

func main() {
	_, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

}
