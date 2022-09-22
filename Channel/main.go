package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.co.in/",
		"https://linkedin.com/",
		"https://amazon.in/",
		"http://udemy.com/",
		"https://golang.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link + " is Okay")
	c <- link
}
