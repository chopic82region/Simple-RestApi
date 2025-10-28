package music

import (
	"errors"
	"time"
)

type ErrMessage struct {
	Error error
	Time  time.Time
}

func NewErrMessage(err string) ErrMessage {
	return ErrMessage{
		Error: errors.New(err),
		Time:  time.Now(),
	}
}

var ErrMusicNotFound = errors.New("Music not founded...")
