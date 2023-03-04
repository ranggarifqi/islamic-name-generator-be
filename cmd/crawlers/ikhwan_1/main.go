package main

import (
	"fmt"
	"strings"

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
	_ = name.NewService(nameRepository)

	// https://www.detik.com/sulsel/berita/d-6529117/1350-nama-bayi-laki-laki-islami-lengkap-beserta-artinya
	c := colly.NewCollector(
		colly.AllowedDomains(
			"detik.com",
			"https://detik.com",
			"www.detik.com",
		),
	)

	var errorArr []error

	c.OnHTML("li", func(h *colly.HTMLElement) {
		class := h.Attr("class")

		if class != "" {
			return
		}

		parent := h.DOM.Parent()
		_, isParentClassExist := parent.Attr("class")
		if isParentClassExist {
			return
		}

		payload, err := constructPayload(h.Text)
		if err != nil {
			errorArr = append(errorArr, err)
		}

		fmt.Printf("payload = %v\n", payload)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.detik.com/sulsel/berita/d-6529117/1350-nama-bayi-laki-laki-islami-lengkap-beserta-artinya/")

	c.Wait()

	for _, err := range errorArr {
		fmt.Printf("Error: %v\n", err)
	}
}

func constructPayload(text string) (*name.Name, error) {
	const gender name.Gender = name.IKHWAN
	nameTypes := [3]name.NameType{name.FIRST_NAME, name.MIDDLE_NAME, name.LAST_NAME}

	trimmed := strings.Trim(text, " ")
	splitted := strings.Split(trimmed, ":")

	if len(splitted) != 2 {
		return nil, fmt.Errorf("error constructing payload: invalid name & meanings: %v", text)
	}

	trimmedName := strings.Trim(splitted[0], " ")
	trimmedMeaning := strings.Trim(splitted[1], " ")

	lowerCasedName := strings.ToLower(trimmedName)
	lowerCasedMeaning := strings.ToLower(trimmedMeaning)

	meaningArr := strings.Split(lowerCasedMeaning, ", ")

	payload := name.Name{
		Name:      lowerCasedName,
		NameTypes: nameTypes[:],
		Gender:    gender,
		Meanings:  meaningArr,
	}

	return &payload, nil
}
