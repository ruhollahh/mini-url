package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	uurl, err := url.Parse("http://localhost:8090")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(uurl.JoinPath("sssss").String())
}
