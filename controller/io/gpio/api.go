package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

func Startup() error {
	return rpio.Open()
}

func PinToggle(pinID int, on bool) {
	fmt.Println(fmt.Sprintf("===> Switching pin %v to %v", pinID, on))
	pin := rpio.Pin(pinID)
	pin.Output()
	if on {
		pin.High()
	} else {
		pin.Low()
	}
}

func PinRead(pinID int) bool {
	pin := rpio.Pin(pinID)
	pin.Input()
	res := pin.Read()
	return res == rpio.High
}

func Close() {
	rpio.Close()
}
