package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b7f45",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {

	for key, value := range m {
		fmt.Println(key, value)
	}
}
