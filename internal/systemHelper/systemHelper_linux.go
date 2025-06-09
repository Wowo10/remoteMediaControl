//go:build linux
// +build linux

package systemHelper

import (
	"log"
	"time"

	"gopkg.in/bendahl/uinput.v1"
)

var device uinput.Keyboard

func init() {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("virtual-media-keyboard"))
	if err != nil {
		log.Fatal(err)
	}

	device = keyboard
}

func SendPlayPause() error {
	device.KeyDown(uinput.KeyPlaypause)
	time.Sleep(50 * time.Millisecond)
	device.KeyUp(uinput.KeyPlaypause)

	return nil
}

func SendNext() error {
	device.KeyDown(uinput.KeyNextsong)
	time.Sleep(50 * time.Millisecond)
	device.KeyUp(uinput.KeyNextsong)

	return nil
}

func SendPrevious() error {
	device.KeyDown(uinput.KeyPrevioussong)
	time.Sleep(50 * time.Millisecond)
	device.KeyUp(uinput.KeyPrevioussong)

	return nil
}

func Dispose() {
	device.Close()
}
