package services

import (
	"fmt"
	"remoteMediaControl/internal/systemHelper"
	"strconv"
)

type Command int

const (
	PlayPause Command = 100
	Next      Command = 101
	Previous  Command = 102
)

func (c Command) String() string {
	switch c {
	case PlayPause:
		return "Play/Pause"
	case Next:
		return "Next"
	case Previous:
		return "Prev"
	default:
		return "Unknown"
	}
}

func HandleWebSocketMessage(messageType int, message []byte) (Command, error) {
	if messageType == 2 {
		return 0, fmt.Errorf("Invalid message type: %d", messageType)
	}

	intCode, err := strconv.Atoi(string(message))
	if err != nil {
		return 0, fmt.Errorf("Invalid message content: %v", err)
	}

	switch Command(intCode) {
	case PlayPause:
		return PlayPause, systemHelper.SendPlayPause()
	case Next:
		return Next, systemHelper.SendNext()
	case Previous:
		return Previous, systemHelper.SendPrevious()
	}

	return 0, fmt.Errorf("Invalid message code: %s", string(message))
}
