package main

import "fmt"

func mapBasics() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors1)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[int]string)
	colors3[10] = "#ffffff"
	fmt.Println(colors3)
	delete(colors3, 10)
	fmt.Println(colors3)

	printMap(colors1)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
