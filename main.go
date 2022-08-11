package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"rsc.io/quote"
)

const myUrl = "https://httpbin.org/"

func main() {
	matrix := make([][]int, 10)
	counter := 0
	for x := 0; x < 10; x++ {

		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter

		}
		fmt.Println(matrix[x])
	}

	fmt.Println(quote.Go())
	MakeRequest()
}
func MakeRequest() {

	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(myUrl+"post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	//	fmt.Println(string(body))
}
