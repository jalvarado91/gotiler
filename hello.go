package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

type TilePos struct {
	center image.Point
	bounds image.Rectangle
}

func calculateDims(targetWidth int, targetHeight int, padding int) (width int, height int) {
	availW := targetWidth - padding*3
	availH := targetHeight - padding*3

	return availW / 2, availH / 2
}

func calculateBounds(imgW int, imgH int, pad int) []image.Point {
	return []image.Point{
		image.Pt(pad, pad),
		image.Pt(pad+imgW, pad),
		image.Pt(pad, pad+imgH+pad),
		image.Pt(pad+imgW+pad, pad+imgH+pad)}
}

func main() {
	targetWidth, targetHeight, padding := 1024, 769, 24
	newWidth, newHeight := calculateDims(targetWidth, targetHeight, padding)

	imgSrc := "testdata/branches.png"
	srcImages := []string{
		"testdata/bold.jpg",
		"testdata/beautiful.jpg",
		"testdata/imagination.jpg",
		"testdata/procrastination.jpg"}

	src, err := imaging.Open(imgSrc)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	canvas := imaging.New(targetWidth, targetHeight, color.White)

	size := src.Bounds().Size()

	fmt.Println(canvas.Bounds())
	fmt.Println(srcImages)
	fmt.Printf("newWidth %d, newHeight %d\n", newWidth, newHeight)
	fmt.Printf("width %d\n", size.X)
	fmt.Printf("height %d\n", size.Y)
}
