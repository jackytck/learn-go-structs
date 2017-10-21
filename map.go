package main

import "fmt"

func mapper() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors1)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[int]string)
	colors3[10] = "#ffffff"
	fmt.Println(colors3)
	delete(colors3, 10)
	fmt.Println(colors3)
}
