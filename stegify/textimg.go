package main

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func generateImageWithText() error {
	// Set the image dimensions
	width := 550
	height := 230

	// Create a new RGBA image with white background
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Over)

	// Set the font face
	face := basicfont.Face7x13

	// Set the color
	col := color.Black

	// Create a new drawer
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  fixedPoint(10, 50), // Adjust the starting position
	}

	// Draw the text on the image
	text := "sueweetielllaaa"
	d.DrawString(text)

	// Save the image as a PNG file
	file, err := os.Create("image_with_text.png")
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func fixedPoint(x, y int) fixed.Point26_6 {
	return fixed.Point26_6{
		X: fixed.Int26_6(x) << 6,
		Y: fixed.Int26_6(y) << 6,
	}
}

func main() {
	err := generateImageWithText()
	if err != nil {
		panic(err)
	}
	fmt.Println("Image generated successfully.")
}
