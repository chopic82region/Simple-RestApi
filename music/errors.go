package music

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrMessage struct {
	Error error
	Time  time.Time
}

func NewErrMessage(err error) *ErrMessage {
	return &ErrMessage{
		Error: err,
		Time:  time.Now(),
	}
}

func (e ErrMessage) ErrToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(nil)
	}

	return string(b)
}

var ErrMusicNotFound = errors.New("Music not founded...")
