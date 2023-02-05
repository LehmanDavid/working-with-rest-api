package main

import (
	"fmt"
	"log"
	"news_api/test/news"
	"time"
)

func main() {
	client, err := news.NewClinet(time.Second * 15)

	if err != nil {
		log.Fatal(err)
	}

	response, err := client.GetNews()
	if err != nil {
		log.Fatal(err)
	}
	for _, asset := range response {
		fmt.Println(asset.PrintInfo())
	}
}
