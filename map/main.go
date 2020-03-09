package main

import "fmt"

func main() {
	colors := map[string] string {
		"red" : "#fff0000",
		"green": "#7450000",
	}

	// var colors map[string]string 

	// colors := make(map[string]string)

	colors["white"] = "#ffffff"
	// delete(colors, "white")

	// map is a reference type
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}