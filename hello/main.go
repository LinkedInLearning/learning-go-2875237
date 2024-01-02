package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Read a text file from the web, Network requests")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type %T\n", resp)
	/*  The response object has a public field simply called body
	and the body will represent everything I need,
	specifically the JSON packet.*/
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body) //now ioutil it's just io i use it according to the course
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	fmt.Print(content)
}
