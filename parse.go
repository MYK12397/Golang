package main

import (
	"fmt"
	"log"
	"net/url"
	)


func main() {
	u, er := url.Parse("http://bing.com/search?q=dotnet")

	if er != nil {
	log.Fatal(er)
	}

	u.Scheme = "https"

	u.Host = "google.com"

	q := u.Query()

	q.Set("q", "golang")

	u.RawQuery = q.Encode()

	fmt.Println(u)
}