package main

import (
	"fmt"

	"github.com/gocolly/colly"
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

	fmt.Println(nameService)

	// https://www.detik.com/sulsel/berita/d-6529117/1350-nama-bayi-laki-laki-islami-lengkap-beserta-artinya
	c := colly.NewCollector(
		colly.AllowedDomains(
			"detik.com",
			"https://detik.com",
			"www.detik.com",
		),
	)

	c.OnHTML("li", func(h *colly.HTMLElement) {
		class := h.Attr("class")

		if class != "" {
			return
		}

		fmt.Printf("name & meaning found: %v\n", h.Text)

		// Split the name & the meaning by character ":"
		// Treat them all as FIRST_NAME, MIDDLE_NAME, and LAST_NAME
		// lower case them all (name & meaning)
		// Set gender = IKHWAN
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.detik.com/sulsel/berita/d-6529117/1350-nama-bayi-laki-laki-islami-lengkap-beserta-artinya/")

	c.Wait()

	// TODO: Crawl the page & get the names, meanings, etc
	// result, err := nameService.UpsertName(name.Name{
	// 	Name:      "Test 1",
	// 	NameTypes: []name.NameType{name.FIRST_NAME},
	// 	Gender:    name.IKHWAN,
	// 	Meanings:  []string{"G ada", "testzzz", "wololo"},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("result = %v\n", result)
}
