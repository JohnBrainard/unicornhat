package main

import (
	"math"
	"time"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/JohnBrainard/unicornhat"
)

/**
This example program was adapted from:
https://github.com/pimoroni/unicorn-hat-hd/blob/master/examples/rainbow.py
*/

func main() {

	var step float64

	display, err := unicornhat.Open(unicornhat.UnicornHD)
	if err != nil {
		panic(err)
	}

	for {
		step++

		for x := 0; x < 16; x++ {
			for y := 0; y < 16; y++ {
				dx := (math.Sin(step/20.0) * 15.0) + 7.0
				dy := (math.Cos(step/15.0) * 15.0) + 7.0
				sc := (math.Cos(step/10.0) * 10.0) + 16.0

				h := math.Sqrt(math.Pow(float64(x)-dx, 2.0)+math.Pow(float64(y)-dy, 2.0)) / sc

				c := colorful.Hsv(math.Mod(h*360.0, 360.0), 1.0, 1.0)

				r := c.R * 255.0
				g := c.G * 255.0
				b := c.B * 255.0

				pixel := unicornhat.NewPixel(byte(r), byte(g), byte(b))
				display.SetPixel(x, y, pixel)
			}
		}

		_ = display.Show()
		time.Sleep(time.Second / 60)
	}
}
