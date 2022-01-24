package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors2 := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
	}
	colors2["white"] = "#ffffff"

	fmt.Println(colors)
	printMap(colors2)
}

func printMap(c map[string]string){
	for color,hex := range c {
		fmt.Println("Hex code for "+color+" is "+hex)
	}
}