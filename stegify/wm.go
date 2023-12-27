package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func generateTransparentImage() error {
	// Set the image dimensions
	width := 550
	height := 230

	// Create a new RGBA image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Set the transparency value (0.0 fully transparent, 1.0 fully opaque)
	transparency := uint8(60)
	transparentColor := color.RGBA{0, 0, 0, transparency}

	// Fill the image with the transparent color
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, transparentColor)
		}
	}

	// Create a PNG file
	file, err := os.Create("transparent_image.png")
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image as PNG and write it to the file
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := generateTransparentImage()
	if err != nil {
		panic(err)
	}
}
