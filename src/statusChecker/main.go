package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- link
		os.Exit(1)
	} else {
		fmt.Println(link, "is up")
		c <- link
	}
}
