package unicornspi

import (
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

type Size int

const (
	UnicornHD Size = 16
	Unicorn   Size = 8
)

// Display encapsulates the SPI connection and pixel buffer.
type Display struct {
	spi    spi.PortCloser
	conn   spi.Conn
	size   int
	buffer []byte
}

// Open opens the SPI port and connection to the Unicorn hat.
func Open(size Size) (*Display, error) {
	_, err := host.Init()
	if err != nil {
		return nil, newOpenError("unable to initialize port host", err)
	}

	port, err := spireg.Open("")
	if err != nil {
		return nil, newOpenError("unable to open default SPI device", err)
	}

	conn, err := port.Connect(90*physic.KiloHertz, spi.Mode0, 8)

	return &Display{
		spi:    port,
		conn:   conn,
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
	packetData := make([]byte, len(d.buffer)+1)
	packetData[0] = 0x72
	copy(packetData[1:], d.buffer)

	packets := []spi.Packet{{
		W:           packetData,
		BitsPerWord: 8,
		KeepCS:      false,
	}}

	if err := d.conn.TxPackets(packets); err != nil {
		return errSendingPackets
	}

	return nil
}

// Close closes the Display and underlying SPI port.
func (d *Display) Close() error {
	if err := d.spi.Close(); err != nil {
		return newCloseError(err)
	}
	return nil
}
