package music

import "errors"

type Music struct {
	Title  string
	Author string

	Downloaded bool
}

func (m Music) IsValidate() error {

	if m.Title == "" {
		return errors.New("Title is empty")
	}
	if m.Author == "" {
		return errors.New("Author is empty")
	}

	return nil
}
