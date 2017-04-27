package unicornhat

import (
	"net"
	"bytes"
)

const (
	SocketPath = "/var/run/unicornd.socket"

	MODE_RGB byte = 0
	MODE_GRB byte = 1

	CMD_SET_BRIGHTNESS byte = 0
	CMD_SET_PIXEL      byte = 1
	CMD_SET_ALL_PIXELS byte = 2
	CMD_SHOW           byte = 3
)

type Hat struct {
	socket net.Conn
	mode   byte
}

func Connect() (*Hat, error) {
	return ConnectToSocket(SocketPath)
}

func ConnectToSocket(path string) (*Hat, error) {
	socket, err := net.Dial("unix", SocketPath)

	if err != nil {
		return nil, err
	}

	return &Hat{socket: socket, mode: MODE_GRB}, nil
}

func (uh *Hat) SetRGBMode() {
	uh.mode = MODE_RGB
}

func (uh *Hat) SetGRBMode() {
	uh.mode = MODE_GRB
}

func (uh *Hat) SetBrightness(brightness byte) error {
	buf := new(bytes.Buffer)

	buf.WriteByte(CMD_SET_BRIGHTNESS)
	buf.WriteByte(brightness)

	_, err := uh.socket.Write(buf.Bytes())
	return err
}

func (uh *Hat) Show() error {
	_, err := uh.socket.Write([]byte{CMD_SHOW})
	return err
}

func (uh *Hat) SetPixel(x byte, y byte, color Color) error {
	buf := new(bytes.Buffer)

	buf.Write([]byte{CMD_SET_PIXEL})
	buf.Write([]byte{x, y})
	buf.Write(color.Bytes(uh.mode))

	_, err := uh.socket.Write(buf.Bytes())
	return err
}

func (uh *Hat) SetAllPixels(pixels [64]Color) error {
	buf := new(bytes.Buffer)

	buf.WriteByte(CMD_SET_ALL_PIXELS)

	for _, color := range pixels {
		buf.Write(color.Bytes(uh.mode))
	}

	_, err := uh.socket.Write(buf.Bytes())
	return err
}

func (uh *Hat) Clear() error {
	return uh.SetAllPixels([64]Color{})
}

func (uh *Hat) Close() error {
	err := uh.socket.Close()
	uh.socket = nil

	return err
}
