package main

import "fmt"

func main() {

	// var colors map[string]string //* same
	// colors := make(map[int]string) //* same [key =string] value=string

	// colors["red"] = "#ff0000"
	// colors[10] = "Ten"

	// delete(colors, 10) //*method ในการลบข้อมูลภายใน map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	//*use func
	printMap(colors)

}

func printMap(c map[string]string) {
	//* for key ,value
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is value ", hex)
	}
}
