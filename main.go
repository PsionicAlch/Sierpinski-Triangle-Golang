package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: [width] [height] [point count]")
		return
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	pixelCount, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	triangle := NewSierpinskiTriangle(width, height, pixelCount)
	triangle.Draw()
	if err := triangle.Save("sierpinski.png"); err != nil {
		log.Fatal(err)
	}
}
