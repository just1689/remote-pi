package io

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
	"github.com/sirupsen/logrus"
)

func Startup() error {
	return rpio.Open()
}

func PinToggle(pinID int, on bool) {
	logrus.Println(fmt.Sprintf("===> Switching pin %v to %v", pinID, on))
	pin := rpio.Pin(pinID)
	pin.Output()
	if on {
		pin.High()
	} else {
		pin.Low()
	}
}

func PinRead(pinID int) (on bool) {
	pin := rpio.Pin(pinID)
	pin.Input()
	res := pin.Read()
	on = res == rpio.High
	logrus.Println(fmt.Sprintf("===> Reading pin %v found it to be: %t", pinID, on))
	return
}

func PinTimeUntilOn(pinID int, timeout time.Duration) (c chan int64) {

	go func() {
		start := time.Now()
		on := false
		pin := rpio.Pin(pinID)
		pin.Input()

		for !on {
			on = pin.Read() == rpio.High
			time.Sleep(1 * time.Millisecond)
			if time.Since(start) > timeout {
				c <- 0
				close(c)
				return
			}
		}

		c <- int64(time.Since(start))
		close(c)
	}()
	return c

}

func Close() {
	rpio.Close()
}
