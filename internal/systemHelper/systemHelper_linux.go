//go:build linux
// +build linux

package systemHelper

import (
	"time"

	"gopkg.in/bendahl/uinput.v1"
)

func SendPlayPause() error {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("virtual-media-keyboard"))
	if err != nil {
		return err
	}
	defer keyboard.Close()

	keyboard.KeyUp(uinput.KeyPlaypause)
	time.Sleep(50 * time.Millisecond)
	keyboard.KeyDown(uinput.KeyPlaypause)

	return nil
}

func SendNext() error {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("virtual-media-keyboard"))
	if err != nil {
		return err
	}
	defer keyboard.Close()

	keyboard.KeyUp(uinput.KeyNextsong)
	time.Sleep(50 * time.Millisecond)
	keyboard.KeyDown(uinput.KeyNextsong)

	return nil
}

func SendPrevious() error {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("virtual-media-keyboard"))
	if err != nil {
		return err
	}
	defer keyboard.Close()

	keyboard.KeyUp(uinput.KeyPrevioussong)
	time.Sleep(50 * time.Millisecond)
	keyboard.KeyDown(uinput.KeyPrevioussong)

	return nil
}
