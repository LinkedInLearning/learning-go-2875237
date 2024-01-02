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
	fmt.Println("Parse JSON-formatted text to make the data structed to use it in my app")
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
	//fmt.Print(content)
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
		err := decoder.Decode((&tour))
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

/*- [David] Thanks for joining me for this course about the Go programming language.
You can learn more about Go from these courses in our library,
including the essential training course where you'll get a deeper dive into
 various aspects of the language, and the course about the Go Standard Library
 where you'll get some experience with real life tasks that you can accomplish
  with various packages and functions that are part of the developer tools.
   To learn more about Visual Studio Code, the IDE that I used throughout the course,
    you can watch this course, which is an introduction to the IDE,
	this one that takes you into a deeper dive and various development tasks,
	and this one that has a series of tips for being more productive using
	Visual Studio Code. These courses and more are available
	to help you as you continue to build your software development skills.*/
