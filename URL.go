package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("We are starting URL in Go language")

	myURL := "http://example.com:8080/path/to/resoiurce?key1=value1&key2=value2"
	fmt.Printf("Type of url is: %T\n", myURL)

	presedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("can't parse URL", presedURL)
		return
	}
	fmt.Printf("Type of URL is: %T\n", presedURL)

	fmt.Println("Scheme: ", presedURL.Scheme)
	fmt.Println("Host: ", presedURL.Host)
	fmt.Println("Path: ", presedURL.Path)
	fmt.Println("RawQuery: ", presedURL.RawQuery)

	//Modifying URL Components

	presedURL.Path = "newpath/to/resource?"
	presedURL.RawQuery = "username=sufiyanshaikh"

	//Constureting a URL String from a URL object

	newurl := presedURL.String()
	fmt.Println("New url is:", newurl)
}
