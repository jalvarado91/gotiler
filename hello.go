package main

import (
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
		image.Pt(pad+pad+imgW, pad),
		image.Pt(pad, pad+imgH+pad),
		image.Pt(pad+imgW+pad, pad+imgH+pad)}
}

func createBoxball(srcImages []string, outPath string, targetWidth int, targetHeight int, padding int) error {
	newWidth, newHeight := calculateDims(targetWidth, targetHeight, padding)

	bounds := calculateBounds(newWidth, newHeight, padding)
	canvas := imaging.New(targetWidth, targetHeight, color.White)

	var subImg image.Image
	var err error
	for idx, imPath := range srcImages {
		subImg, err = imaging.Open(imPath)
		subImg = imaging.Fill(subImg, newWidth, newHeight, imaging.Center, imaging.Lanczos)
		if err != nil {
			return err
		}
		canvas = imaging.Paste(canvas, subImg, bounds[idx])
	}

	err = imaging.Save(canvas, outPath)
	return err
}

func main() {
	targetWidth, targetHeight, padding := 1024, 768, 16
	srcImages := []string{
		"testdata/bold.jpg",
		"testdata/beautiful.jpg",
		"testdata/imagination.jpg",
		"testdata/procrastination.jpg"}

	err := createBoxball(srcImages, "testdata/out_tile.jpg", targetWidth, targetHeight, padding)
	if err != nil {
		log.Fatalf("failed to save create boxball layou: %v", err)
	}
}
