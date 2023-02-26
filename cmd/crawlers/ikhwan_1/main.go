package main

import (
	"fmt"

	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

func main() {
	mongoDBClient, ctx, cancel, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

	defer cancel()
	defer mongodb.Disconnect(mongoDBClient, ctx)

	mongoDB := mongoDBClient.Database("islamic-name-generator")

	nameRepository := name.NewMongoRepository(ctx, mongoDB)
	nameService := name.NewService(nameRepository)

	// TODO: Crawl the page & get the names, meanings, etc
	result, err := nameService.UpsertName(name.Name{
		Name:      "Test 1",
		NameTypes: []name.NameType{name.FIRST_NAME},
		Gender:    name.IKHWAN,
		Meanings:  []string{"G ada", "testzzz", "wololo"},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("result = %v\n", result)
}
