package services

import (
	"fmt"
	"remoteMediaControl/internal/systemHelper"
	"strconv"
)

func HandleWebSocketMessage(messageType int, message []byte) error {
	if messageType == 1 {
		return fmt.Errorf("Invalid message type: %d", messageType)
	}

	intCode, err := strconv.Atoi(string(message))
	if err != nil {
		return fmt.Errorf("Invalid message content: %v", err)
	}

	switch intCode {

	case 100:
		return systemHelper.SendPlayPause()
	case 101:
		return systemHelper.SendNext()
	case 102:
		return systemHelper.SendPrevious()
	}

	return fmt.Errorf("Invalid message content: %s", string(message))
}
