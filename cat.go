package main

import (
	"io"
	"log"
	"os"
)

func cat(s string) {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatal(err)
	}
}
