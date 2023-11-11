package main

import "fmt"

func main() {

	//Different ways of initializing map

	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"white": "#ffffff",
	// }

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	fmt.Println(colors["black"]) //Use square braces to access a value

	delete(colors, "black")

	fmt.Println(colors)

}
