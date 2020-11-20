package main

import (
	"fmt"
	"net/http"
	"time"
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
		go checklink(link, c)
	}

	// for {
	// 	go checklink(<-c, c)
	// }

	for l := range c {

		go func(link string) {
			time.Sleep(15 * time.Second)
			checklink(link, c)
		}(l)
	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
