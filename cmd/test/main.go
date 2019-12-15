package main

import (
	"log"
	"time"

	"github.com/JohnBrainard/unicornhat/unicornspi"
)

func main() {
	hat, err := unicornspi.Open(unicornspi.UnicornHD)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := hat.Close(); err != nil {
			panic(err)
		}
	}()

	pixelTemplates := []unicornspi.Pixel{
		unicornspi.NewPixel(255, 0, 0),
		unicornspi.NewPixel(0, 255, 0),
		unicornspi.NewPixel(0, 0, 255),
	}

	for _, template := range pixelTemplates {
		pixels := make([]unicornspi.Pixel, 16*16)

		for i, _ := range pixels {
			x := i % 16
			y := i / 16

			if y < 4 && (x == 7 || x == 8) {
				pixels[i] = template.Invert()
			} else {
				pixels[i] = template
			}
		}
		hat.SetPixels(pixels)
		_ = hat.Show()
		time.Sleep(time.Second * 10)
	}

	hat.Clear()
	if err := hat.Show(); err != nil {
		panic(err)
	}
}
