package main

import (
	"flag"
	"fmt"
)

type Color int

const (
	Cyan  Color = 6
	White Color = 7
)

const (
	tlCorner = 0x250c
	trCorner = 0x2510
	blCorner = 0x2514
	brCorner = 0x2518
	hLine    = 0x2500
	vLine    = 0x2502
)

type CharSet struct {
	title      string
	start, end int
}

func writeChar(color Color, char int) {
	if (char >> 8) == 0 {
		fmt.Printf("00%x:\u001b[38;5;%dm %c \u001b[0m", char, color, char)
	} else {
		fmt.Printf("%x:\u001b[38;5;%dm %c \u001b[0m", char, color, char)
	}
}

func (c CharSet) draw() {
	width := 129
	fmt.Printf("%c", tlCorner)
	fmt.Printf("%s", c.title)
	for i := 0; i < (width - len(c.title)); i++ {
		fmt.Printf("%c", hLine)
	}
	fmt.Printf("%c\n", trCorner)

	for y := c.start; y <= c.end; y += 0x0010 {
		fmt.Printf("%c ", vLine)

		for x := 0x00; x < 0x10; x++ {
			switch y + x {
			case 0x007f:
				fmt.Printf("00%x: \u001b[38;5;6m%c \u001b[0m", y+x, 0x2421)
			case 0x25fd, 0x25fe:
				fmt.Printf("%x:\u001b[38;5;6m%c \u001b[0m", y+x, y+x)
			default:
				writeChar(Cyan, y+x)
			}
		}
		fmt.Printf("%c\n", vLine)
	}

	fmt.Printf("%c", blCorner)
	for i := 0; i < width; i++ {
		fmt.Printf("%c", hLine)
	}
	fmt.Printf("%c\n", brCorner)
}

func ctrlCodes() {
	code := map[int]int{
		0x0001: 0x2191,
		0x0002: 0x2193,
		0x0003: 0x2192,
		0x0004: 0x2190,
		0x0008: 0x232b,
		0x0009: 0x21e5,
		0x000d: 0x23ce,
		0x001b: 0x21b5,
	}

	fmt.Printf("%c", 0x250c)
	fmt.Printf("%s", "CTRL Codes")
	for i := 0; i < 119; i++ {
		fmt.Printf("%c", 0x2500)
	}
	fmt.Printf("%c\n", 0x2510)
	for y := 0x0001; y < 0x0020; y += 0x0010 {
		fmt.Printf("%c ", 0x2502)
		for x := 0x00; x < 0x10; x++ {
			val, ok := code[y+x]
			if ok {
				fmt.Printf("000%x: \u001b[38;5;6m%c", y+x, val)
			} else if y+x < 0x0010 {
				fmt.Printf("000%x:\u001b[38;5;6m%s%c", y+x, "^", y+x+0x0040)
			} else if y+x > 0x001a {
				fmt.Printf("%s\u001b[38;5;6m%c", "      ", 0x0020)
			} else {
				fmt.Printf("00%x:\u001b[38;5;6m%s%c", y+x, "^", y+x+0x0040)
			}
			if y+x == 0x001b {
				fmt.Printf("\u001b[0m")
			} else {
				fmt.Printf(" \u001b[0m")
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
	box := CharSet{"Box Drawing", 0x2500, 0x257F}
	block := CharSet{"Block Drawing", 0x2580, 0x259F}
	shape := CharSet{"Shape Drawing", 0x25A0, 0x25FF}
	braille := CharSet{"Braille", 0x2800, 0x28ff}
	ascii := CharSet{"ASCII", 0x0020, 0x007F}
	charSets := []CharSet{box, block, shape, braille, ascii}

	// Add command line flags to show specific tables
	boxFlag := flag.Bool("box", false, "displays box drawing characters")
	blockFlag := flag.Bool("block", false, "displays block drawing characters")
	shapeFlag := flag.Bool("shape", false, "displays shape drawing characters")
	brailleFlag := flag.Bool("braille", false, "displays braille characters")
	asciiFlag := flag.Bool("ascii", false, "displays ascii characters")
	ctrlFlag := flag.Bool("ctrl", false, "displays control codes")
	allFlag := flag.Bool("all", false, "displays all tables")

	flag.Parse()

	if *boxFlag {
		box.draw()
	}
	if *blockFlag {
		block.draw()
	}
	if *shapeFlag {
		shape.draw()
	}
	if *brailleFlag {
		braille.draw()
	}
	if *asciiFlag {
		ascii.draw()
	}
	if *ctrlFlag {
		ctrlCodes()
	}
	if *allFlag || !(*boxFlag || *blockFlag || *shapeFlag || *brailleFlag || *asciiFlag || *ctrlFlag) {
		for _, i := range charSets {
			i.draw()
		}
		ctrlCodes()
	}
}
