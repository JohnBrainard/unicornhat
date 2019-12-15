package unicornhat

import "github.com/JohnBrainard/unicornhat/spi"

type Size int

const (
	UnicornHD Size = 16
	Unicorn   Size = 8
)

// Display encapsulates the SPI connection and pixel buffer.
type Display struct {
	driver Driver
	size   int
	buffer []byte
}

// Open opens the SPI port and connection to the Unicorn hat.
func Open(size Size) (*Display, error) {
	var driver Driver
	var err error

	switch size {
	case UnicornHD:
		driver, err = spi.New()
		if err != nil {
			return nil, err
		}

	default:
		return nil, errIncompatibleDevice
	}

	return &Display{
		driver: driver,
		size:   int(size),
		buffer: make([]byte, size*size*3),
	}, nil
}

// SetPixel sets the RGB data at the X,Y coordinates in the buffer.
func (d *Display) SetPixel(x, y int, pixel Pixel) error {
	if x > d.size || y > d.size {
		panic(errBufferOverflow)
	}

	i := x*d.size + y
	return d.SetPixelAt(i, pixel)
}

// SetPixelAt sets the RGB data at the provided position in the buffer.
func (d *Display) SetPixelAt(pos int, pixel Pixel) error {
	if pos < 0 || pos+3 > len(d.buffer) {
		panic(errBufferOverflow)
	}

	d.buffer[3*pos] = pixel.r
	d.buffer[3*pos+1] = pixel.g
	d.buffer[3*pos+2] = pixel.b

	return nil
}

// SetPixels sets the pixel buffer to the provided pixels.
func (d *Display) SetPixels(pixels []Pixel) {
	for i, pixel := range pixels {
		if err := d.SetPixelAt(i, pixel); err != nil {
			panic(errBufferOverflow)
		}
	}
}

// Clear clears the pixel buffer.
func (d *Display) Clear() {
	d.buffer = make([]byte, d.size*d.size*3)
}

// Show sends the pixel buffer to the Unicorn hat.
func (d *Display) Show() error {
	return d.driver.Render(d.buffer)
}

// Close closes the Display and underlying SPI port.
func (d *Display) Close() error {
	return d.driver.Close()
}
