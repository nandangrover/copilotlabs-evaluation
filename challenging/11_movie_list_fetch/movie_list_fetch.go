package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://ghibliapi.herokuapp.com/films")
	if err != nil {
		fmt.Println("Could not send request to API:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Expected status code 200 but received", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read the response body:", err)
		return
	}
	fmt.Println("Processing our list of movies")
	fmt.Println(string(body))
}
