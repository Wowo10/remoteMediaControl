//go:build windows
// +build windows

package systemHelper

import (
	"github.com/micmonay/keybd_event"
)

func SendPlayPause() error {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_PLAY)
	return kb.Launching()
}

func SendNext() error {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_MEDIA_NEXT_TRACK)
	return kb.Launching()
}

func SendPrevious() error {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_MEDIA_PREV_TRACK)
	return kb.Launching()
}
