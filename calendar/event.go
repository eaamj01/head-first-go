package calendar

import (
	"errors"
	"unicode/utf8"
)

// Event struct that holds and event and date
type Event struct {
	title string
	Date
}

// Title get title value
func (e *Event) Title() string {
	return e.title
}

// SetTitle Validate that incoming title is no longer than 30 characters and then set Event.title
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
