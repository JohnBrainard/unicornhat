package unicornspi

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

type unicornError string

var (
	errBufferOverflow unicornError = "buffer overflow"
	errSendingPackets unicornError = "sending packets"
)

func (e unicornError) Error() string {
	return string(e)
}

func InitSPI() {
	fmt.Println("Initializing SPI")
	_, err := host.Init()
	if err != nil {
		log.Fatalf("failed initializing SPI: %v", err)
	}
	spiPort, err := spireg.Open("0")
	if err != nil {
		log.Fatalf("unable to open spi: %v", err)
	}

	defer func() {
		if err := spiPort.Close(); err != nil {
			log.Fatal("error closing SPI", err)
		}
	}()
}
