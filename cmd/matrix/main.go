package main

import (
	"math/rand"
	"time"

	"github.com/JohnBrainard/unicornhat"
)

/**
This example program was adapted from:
https://github.com/pimoroni/unicorn-hat-hd/blob/master/examples/matrix-hd.py
*/

type person struct {
	x, y int
}

func (p *person) moveDown() {
	p.y--
}

func NewPerson() *person {
	return &person{
		x: rand.Intn(15),
		y: 15,
	}
}

func main() {
	wordRGB := []unicornhat.Pixel{
		unicornhat.NewPixel(154, 173, 154),
		unicornhat.NewPixel(0, 255, 0),
		unicornhat.NewPixel(0, 235, 0),
		unicornhat.NewPixel(0, 220, 0),
		unicornhat.NewPixel(0, 185, 0),
		unicornhat.NewPixel(0, 165, 0),
		unicornhat.NewPixel(0, 128, 0),
		unicornhat.NewPixel(0, 0, 0),
		unicornhat.NewPixel(154, 173, 154),
		unicornhat.NewPixel(0, 145, 0),
		unicornhat.NewPixel(0, 125, 0),
		unicornhat.NewPixel(0, 100, 0),
		unicornhat.NewPixel(0, 80, 0),
		unicornhat.NewPixel(0, 60, 0),
		unicornhat.NewPixel(0, 40, 0),
		unicornhat.NewPixel(0, 0, 0),
	}

	display, err := unicornhat.Open(unicornhat.UnicornHD)
	if err != nil {
		panic(err)
	}
	display.Rotate(270)
	step := 0
	population := []*person{
		NewPerson(),
	}

	for {
		for _, person := range population {
			// fmt.Printf("person x,y = %d,%d\n", person.x, person.y)
			y := person.y
			for _, pixel := range wordRGB {
				if (y <= 15) && (y >= 0) {
					if err := display.SetPixel(person.x, y, pixel); err != nil {
						panic(err)
					}
				}
				y++
			}
			person.moveDown()
		}

		err = display.Show()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second / 10)

		step++

		if step%5 == 0 || step%7 == 0 {
			population = append(population, NewPerson())
		}

		if len(population) > 100 {
			population = population[len(population)-100:]
		}
	}
}
