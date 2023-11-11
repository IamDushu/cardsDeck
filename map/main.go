package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"green": "#4bf745",
	}

	printMap(colors)
	//Hex code for red is #ff0000
	//Hex code for white is #ffffff
	//Hex code for green is #4bf745

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
