package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type SierpinskiTriangle struct {
	Img            *image.Gray
	PointCount     int
	Width          int
	Height         int
	OriginalPoints [3]image.Point
}

// NewSierpinskiTriangle creates a new instance of the SierpinskiTriangle. This function
// takes the width and height of the image along with how many points you want to set on
// the triangle.
func NewSierpinskiTriangle(width, height, pointCount int) *SierpinskiTriangle {
	// Create a new image.
	img := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	// Set all the pixels to black.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.Black)
		}
	}

	// Set the initial 3 points.
	var point1, point2, point3 image.Point

	// Top Point.
	point1.X = width / 2
	point1.Y = 5

	// Bottom Left Point.
	point2.X = 5
	point2.Y = height - 5

	// Bottom Right Point.
	point3.X = width - 5
	point3.Y = height - 5

	points := [3]image.Point{point1, point2, point3}

	return &SierpinskiTriangle{
		Img:            img,
		PointCount:     pointCount,
		Width:          width,
		Height:         height,
		OriginalPoints: points,
	}
}

type void struct{}

var member void

func (t *SierpinskiTriangle) Draw() {
	// Create a set so that we only save original points.
	pixelSet := make(map[image.Point]void)

	for _, point := range t.OriginalPoints {
		pixelSet[point] = member
	}

	fmt.Println("Finished setting the initial 3 points.")

	rp := image.Point{X: rand.Intn(t.Width), Y: rand.Intn(t.Height)}
	for !pointInTriangle(rp, t.OriginalPoints[0], t.OriginalPoints[1], t.OriginalPoints[2]) {
		rp = image.Point{X: rand.Intn(t.Width), Y: rand.Intn(t.Height)}
	}

	pixelSet[rp] = member

	for len(pixelSet) < t.PointCount+3 {
		rand.Seed(time.Now().UnixNano())

		ogp := t.OriginalPoints[rand.Intn(3)]

		mp := midpoint(ogp, rp)
		pixelSet[mp] = member

		fmt.Printf("%d/%d points set.\n", len(pixelSet)-3, t.PointCount)

		rp = mp
	}

	fmt.Println("Finished setting the other points.")

	for point, _ := range pixelSet {
		t.Img.Set(point.X, point.Y, color.White)
	}
}

func (t *SierpinskiTriangle) Save(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	err = png.Encode(file, t.Img)
	if err != nil {
		return err
	}

	return nil
}

// I got this function from https://stackoverflow.com/a/2049593
func sign(p1, p2, p3 image.Point) int {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}

// I got this function from https://stackoverflow.com/a/2049593
func pointInTriangle(point, p1, p2, p3 image.Point) bool {
	var d1, d2, d3 int
	var hasNeg, hasPos bool

	d1 = sign(point, p1, p2)
	d2 = sign(point, p2, p3)
	d3 = sign(point, p3, p1)

	hasNeg = (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos = (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(hasNeg && hasPos)
}

func midpoint(p1, p2 image.Point) image.Point {
	return image.Point{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}
}
