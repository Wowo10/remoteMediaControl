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

func SendVolUp() error {
	device.KeyDown(uinput.KeyVolumeup)
	time.Sleep(50 * time.Millisecond)
	device.KeyUp(uinput.KeyVolumeup)

	return nil
}

func SendVolDown() error {
	device.KeyDown(uinput.KeyVolumedown)
	time.Sleep(50 * time.Millisecond)
	device.KeyUp(uinput.KeyVolumedown)

	return nil
}

func Dispose() {
	device.Close()
}
