package main

func main() {

	// var colors map[string]string
	// colors = make(map[string]string)
	// colors["white"] = "#FFFFFF"

	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#EFBC00",
		// "Green": "#FF0011",
		"green": "#AA0099",
	}

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println("Color: ", color, "Hex: ", hex)
	}
}
