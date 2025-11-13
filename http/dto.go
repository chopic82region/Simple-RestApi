package http

import (
	"encoding/json"
	"errors"
	"time"
)

type MusicDTO struct {
	Name   string
	Author string
}

func (m MusicDTO) IsValidate() error {

	if m.Name == "" {
		return errors.New("Title is empty")
	}
	if m.Author == "" {
		return errors.New("Author is empty")
	}

	return nil
}

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
		panic(err)
	}

	return string(b)
}
