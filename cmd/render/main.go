package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/JohnBrainard/unicornhat"
)

func main() {
	display, err := unicornhat.Open(unicornhat.UnicornHD)
	if err != nil {
		panic(err)
	}

	display.Rotate(90)

	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	imagePath := os.Args[1]
	var img image.Image

	switch path.Ext(imagePath) {
	case ".png":
		img = loadPng(imagePath)
	}

	displayImage(img, display)
}

func loadPng(imagePath string) image.Image {
	reader, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("unable to open image: %s", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Fatalf("failed to close image: %v", err)
		}
	}()

	pngImage, err := png.Decode(reader)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	return pngImage
}

func displayImage(img image.Image, display *unicornhat.Display) {
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			pixel := unicornhat.NewPixel(
				byte(r),
				byte(g),
				byte(b))

			if err := display.SetPixel(x, y, pixel); err != nil {
				panic(err)
			}
		}
	}

	if err := display.Show(); err != nil {
		log.Fatalf("unable to show pixels: %v", err)
	}
}

func printUsage() {
	fmt.Printf(`Usage:
%s <png>

<png> PNG image to render.
`, path.Base(os.Args[0]))
}
