package services

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func FetchDataFromAPI() {
	client := resty.New()
	resp, err := client.R().
		Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Printf("Error occurred while fetching data: %v", err)
	}

	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Response Body:", resp.String())
}
