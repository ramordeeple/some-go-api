package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		pic[y] = row
	}

	return pic
}

func printPic(pic [][]uint8) {
	for _, row := range pic { 
		for _, value := range row {
			fmt.Print(value, "")
		}
		fmt.Println()
	}
}

func main() {
	image := Pic(10, 5)
	printPic(image)
}
