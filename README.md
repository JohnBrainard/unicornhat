# UnicornHat

UnicornHat is a client library for [unicornd][unicornd]

```sh
$ go get github.com/JohnBrainard/unicornhat
```

## Example

```go
package main

import (
	"log"
	"math/rand"
	"time"
	"flag"
	"github.com/JohnBrainard/UnicornHat"
)

func main() {
	r := flag.Uint("r", 0, "red")
	g := flag.Uint("g", 0, "green")
	b := flag.Uint("b", 0, "blue")

	clear := flag.Bool("clear", false, "clear display")
	random := flag.Bool("random", false, "random")

	flag.Parse()

	client, err := unicornhat.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if *random {
		rand.Seed(time.Now().UnixNano())

		pixels := [64]unicornhat.Color{}

		for i := range pixels {
			r := byte(rand.Float32() * 256)
			g := byte(rand.Float32() * 256)
			b := byte(rand.Float32() * 256)

			pixels[i] = unicornhat.ColorNew(r, g, b)
			client.SetAllPixels(pixels)
		}
	} else if *clear {
		client.Clear()
	} else {
		pixels := [64]unicornhat.Color{}
		for i := range pixels {
			pixels[i] = unicornhat.ColorNew(
				byte(*r),
				byte(*g),
				byte(*b),
			)
		}
		client.SetAllPixels(pixels)
	}

	client.Show()
}

```

[unicornd]:https://github.com/pimoroni/unicorn-hat/tree/master/library_c/unicornd