package main

import (
	"fmt"
)

func main() {
	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#fgciud",
	// }

	// var color map[string]string

	color := make(map[string]string)

	color["white"] = "#ffffff"
	color["red"] = "#ff0000"
	color["green"] = "#fgciud"

	fmt.Printf("%v", color)
	fmt.Println()

	// delete(color, "white")

	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color : " + color + " hex : " + hex)
	}
}
