package main

import (
	"fmt"
	"net/http"
	"os"
)

func getRequest() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	buf := make([]byte, 99999)
	n, _ := resp.Body.Read(buf)
	fmt.Println(string(buf))
	fmt.Println(n)
	resp.Body.Close()
}
