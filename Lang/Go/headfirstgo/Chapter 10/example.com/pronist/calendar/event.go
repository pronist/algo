package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e Event) Title() string {
	return e.title
}

func (e *Event) SetTitle() error {
	if utf8.RuneCountInString(e.title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
