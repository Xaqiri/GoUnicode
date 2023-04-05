package main

import (
	"flag"
	"fmt"
	"strings"
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

func helpDrawing() {
	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "Help")
	for i := 0; i < 62; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	fmt.Printf("%c --help: Displays this help table%s%c\n", 0x2502, strings.Repeat(" ", 33), 0x2502)
	fmt.Printf("%c --block: Displays the unicode table for block drawing characters %c\n", 0x2502, 0x2502)
	fmt.Printf("%c --box: Displays the unicode table for box drawing characters%s%c\n", 0x2502, strings.Repeat(" ", 5), 0x2502)
	fmt.Printf("%c --ascii: Displays the unicode table for ascii characters%s%c\n", 0x2502, strings.Repeat(" ", 9), 0x2502)

	fmt.Printf("%c", 0x2514)
	for i := 0; i < 66; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2518)

}

func main() {
	// Add command line flags to show specific tables
	// help := flag.Bool("help", false, "displays help")
	box := flag.Bool("box", false, "displays box drawing characters")
	block := flag.Bool("block", false, "displays block drawing characters")
	shape := flag.Bool("shape", false, "displays shape drawing characters")
	ascii := flag.Bool("ascii", false, "displays ascii characters")
	all := flag.Bool("all", false, "displays all tables")

	flag.Parse()
	// if *help {
	// helpDrawing()
	// } else {
	if *box {
		boxDrawing()
	}
	if *block {
		blockDrawing()
	}
	if *shape {
		shapeDrawing()
	}
	if *ascii {
		asciiDrawing()
	}
	if *all || !(*box || *block || *shape || *ascii) {
		boxDrawing()
		blockDrawing()
		shapeDrawing()
		asciiDrawing()
	}
}
