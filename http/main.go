package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("https://www.google.com/search?q=free")

	// body := make([]byte, 999999999)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// resp.Body.Read(body)
	// fmt.Println(string(body))

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
