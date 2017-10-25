package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func getRequest() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// buf := make([]byte, 99999)
	// n, _ := resp.Body.Read(buf)
	// fmt.Println(string(buf))
	// fmt.Println(n)

	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
