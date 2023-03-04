package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

func main() {
	mongoDBClient, ctx, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

	mongoDB := mongoDBClient.Database("islamic-name-generator")

	nameRepository := name.NewMongoRepository(ctx, mongoDB)
	nameService := name.NewService(nameRepository)

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
			return
		}

		insertedName, err := nameService.UpsertName(*payload)
		if err != nil {
			errorArr = append(errorArr, err)
			return
		}
		fmt.Printf("%v upserted successfully; %v\n", insertedName.Name, *insertedName)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.detik.com/sulsel/berita/d-6529117/1350-nama-bayi-laki-laki-islami-lengkap-beserta-artinya/")

	c.Wait()

	defer mongodb.Disconnect(mongoDBClient, ctx)

	if len(errorArr) > 0 {
		writeErrorsIntoFile("./", errorArr)
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

func writeErrorsIntoFile(dirPath string, errors []error) {
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("error_%v.log", timestamp)
	filePath := path.Join(dirPath, fileName)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, error := range errors {
		_, err = f.WriteString(fmt.Sprintf("%v\n", error.Error()))
		if err != nil {
			panic(err)
		}
	}

	f.Sync()
}
