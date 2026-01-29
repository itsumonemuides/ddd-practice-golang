package todo

import "errors"

type Title struct {
	value string
}

// NewTitle create a new Title with validation
func NewTitle(title string) (Title, error) {
	if title == "" {
		return Title{}, errors.New("title cannnot be empty")
	}
	if len(title) > 100 {
		return Title{}, errors.New("title must be 100 characters or less")
	}
	return Title{value: title}, nil
}

func (t Title) String() string {
	return t.value
}
