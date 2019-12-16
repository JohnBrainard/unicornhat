package unicornhat

import "fmt"

type Pixel struct {
	r, g, b byte
}

func NewPixel(r, g, b byte) Pixel {
	return Pixel{
		r: r,
		g: g,
		b: b,
	}
}

func (p Pixel) Bytes() []byte {
	return []byte{p.r, p.g, p.b}
}

func (p Pixel) Invert() Pixel {
	return Pixel{
		r: 255 - p.r,
		g: 255 - p.g,
		b: 255 - p.b,
	}
}

func (p Pixel) String() string {
	return fmt.Sprintf("Pixel(R: %d, G: %d, B:%d)", p.r, p.g, p.b)
}
