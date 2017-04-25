package unicornhat

import (
	"net"
	"bytes"
)

const SocketPath = "/var/run/unicornd.socket"

const CMD_SET_BRIGHTNESS byte = 0
const CMD_SET_PIXEL byte = 1
const CMD_SET_ALL_PIXELS byte = 2
const CMD_SHOW byte = 3

type UnicornHat struct {
	socket net.Conn
}

func Connect() (*UnicornHat, error) {
	return ConnectToSocket(SocketPath)
}

func ConnectToSocket(path string) (*UnicornHat, error) {
	socket, err := net.Dial("unix", SocketPath)

	if err != nil {
		return nil, err
	}

	return &UnicornHat{socket: socket}, nil
}

func (uh *UnicornHat) SetBrightness(brightness byte) {
	buf := new(bytes.Buffer)

	buf.WriteByte(CMD_SET_BRIGHTNESS)
	buf.WriteByte(brightness)

	uh.socket.Write(buf.Bytes())
}

func (uh *UnicornHat) Show() {
	uh.socket.Write([]byte{CMD_SHOW})
}

func (uh *UnicornHat) SetPixel(x byte, y byte, color Color) {
	buf := new(bytes.Buffer)

	buf.Write([]byte{CMD_SET_PIXEL})
	buf.Write([]byte{x, y})
	buf.Write(color.Bytes())

	uh.socket.Write(buf.Bytes())
}

func (uh *UnicornHat) SetAllPixels(pixels [64]Color) {
	buf := new(bytes.Buffer)

	buf.WriteByte(CMD_SET_ALL_PIXELS)

	for _, color := range pixels {
		buf.Write(color.Bytes())
	}

	uh.socket.Write(buf.Bytes())
}

func (uh *UnicornHat) Clear() {
	uh.SetAllPixels([64]Color{})
}
