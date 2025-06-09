//go:build windows
// +build windows

package systemHelper

import (
	"log"

	"github.com/micmonay/keybd_event"
)

var device keybd_event.KeyBonding

func init() {
	kb, err := keybd_event.NewKeyBonding()

	if err != nil {
		log.Fatal(err)
	}

	device = kb
}

func SendPlayPause() error {
	device.SetKeys(keybd_event.VK_PLAY)
	return device.Launching()
}

func SendNext() error {
	device.SetKeys(keybd_event.VK_MEDIA_NEXT_TRACK)
	return device.Launching()
}

func SendPrevious() error {
	device.SetKeys(keybd_event.VK_MEDIA_PREV_TRACK)
	return device.Launching()
}

func SendVolUp() error {
	device.SetKeys(keybd_event.VK_VOLUME_UP)
	return device.Launching()
}

func SendVolDown() error {
	device.SetKeys(keybd_event.VK_VOLUME_DOWN)
	return device.Launching()
}

func Dispose() {
	// do nothing
}
