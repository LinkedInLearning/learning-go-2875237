package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	// fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
