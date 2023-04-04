package main

import (
	"fmt"
)

func boxDrawing() {
	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "Box Drawing")
	for i := 0; i < 118; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	for y := 0x2500; y < 0x2570; y += 0x0010 {
		fmt.Printf("%c ", 0x2502)

		for x := 0x00; x < 0x10; x++ {
			fmt.Printf("%x: \u001b[38;5;6m%c \u001b[0m", y+x, y+x)
		}
		fmt.Printf("%c\n", 0x2502)
	}
	fmt.Printf("%c", 0x2514)
	for i := 0; i < 129; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2518)
}

func blockDrawing() {
	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "Block Drawing")
	for i := 0; i < 116; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	for y := 0x2580; y <= 0x2590; y += 0x0010 {
		fmt.Printf("%c ", 0x2502)

		for x := 0x00; x < 0x10; x++ {
			fmt.Printf("%x: \u001b[38;5;6m%c \u001b[0m", y+x, y+x)
		}
		fmt.Printf("%c\n", 0x2502)
	}
	fmt.Printf("%c", 0x2514)
	for i := 0; i < 129; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2518)
}

func shapeDrawing() {
	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "Shape Drawing")
	for i := 0; i < 116; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	for y := 0x25A0; y <= 0x25F0; y += 0x0010 {
		fmt.Printf("%c ", 0x2502)
		for x := 0x00; x < 0x10; x++ {
			if y+x == 0x25fd || y+x == 0x25fe {
				fmt.Printf("%x:\u001b[38;5;6m%c \u001b[0m", y+x, y+x)

			} else {
				fmt.Printf("%x: \u001b[38;5;6m%c \u001b[0m", y+x, y+x)
			}
		}
		fmt.Printf("%c\n", 0x2502)
	}
	fmt.Printf("%c", 0x2514)
	for i := 0; i < 129; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2518)
}

func asciiDrawing() {
	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "ASCII")
	for i := 0; i < 124; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	for y := 0x0021; y < 0x0080; y += 0x0010 {
		fmt.Printf("%c ", 0x2502)
		for x := 0x00; x < 0x10; x++ {
			if y+x == 0x007f || y+x == 0x0080 {
				fmt.Printf("00%x: \u001b[38;5;6m%c \u001b[0m", y+x, 0x0020)
			} else {
				fmt.Printf("00%x: \u001b[38;5;6m%c \u001b[0m", y+x, y+x)
			}
		}
		fmt.Printf("%c\n", 0x2502)
	}
	fmt.Printf("%c", 0x2514)
	for i := 0; i < 129; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2518)

}

func main() {
	boxDrawing()
	blockDrawing()
	shapeDrawing()
	asciiDrawing()
}
