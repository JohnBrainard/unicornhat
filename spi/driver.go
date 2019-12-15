package spi

import (
	"periph.io/x/periph/conn/physic"
	spi2 "periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

type spiDriverError string

func (e spiDriverError) Error() string {
	return string(e)
}

const (
	errSendingPackets     spiDriverError = "sending packets"
	errPortInitialization spiDriverError = "unable to initialize port host"
	errOpenSPIDevice      spiDriverError = "unable to open SPI device"
)

type spiDriver struct {
	spi  spi2.PortCloser
	conn spi2.Conn
}

func New() (*spiDriver, error) {
	_, err := host.Init()
	if err != nil {
		return nil, errPortInitialization
	}

	port, err := spireg.Open("")
	if err != nil {
		return nil, errOpenSPIDevice
	}

	conn, err := port.Connect(90*physic.KiloHertz, spi2.Mode0, 8)

	return &spiDriver{
		spi:  port,
		conn: conn,
	}, nil
}

func (s *spiDriver) Render(buffer []byte) error {
	packetData := make([]byte, len(buffer)+1)
	packetData[0] = 0x72
	copy(packetData[1:], buffer)

	packets := []spi2.Packet{{
		W:           packetData,
		BitsPerWord: 8,
		KeepCS:      false,
	}}

	if err := s.conn.TxPackets(packets); err != nil {
		return errSendingPackets
	}

	return nil
}

func (s *spiDriver) Close() error {
	if err := s.spi.Close(); err != nil {
		return err
	}
	return nil
}
