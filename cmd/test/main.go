package main

import (
	"log"
	"time"

	"github.com/JohnBrainard/unicornhat"
)

func main() {
	hat, err := unicornhat.Open(unicornhat.UnicornHD)
	if err != nil {
		log.Panic(err)
	}

	rotations := []int{
		0,
		90,
		180,
		270,
	}

	// hat.Rotate(90)

	defer func() {
		if err := hat.Close(); err != nil {
			panic(err)
		}
	}()

	pixelTemplates := []unicornhat.Pixel{
		unicornhat.NewPixel(255, 0, 0),
		unicornhat.NewPixel(0, 255, 0),
		unicornhat.NewPixel(0, 0, 255),
		unicornhat.NewPixel(255, 255, 0),
		unicornhat.NewPixel(255, 0, 255),
		unicornhat.NewPixel(0, 255, 255),
	}

	for _, rotation := range rotations {
		hat.Rotate(rotation)
		for _, template := range pixelTemplates {
			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					if y < 4 && (x == 7 || x == 8) {
						hat.SetPixel(x, y, template.Invert())
					} else {
						hat.SetPixel(x, y, template)
					}
				}
			}
			_ = hat.Show()
			time.Sleep(time.Second / 3)
		}
	}

	hat.Clear()
	if err := hat.Show(); err != nil {
		panic(err)
	}
}
