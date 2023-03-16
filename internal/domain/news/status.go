package news

import (
	"fmt"
)

type Status struct {
	value string
}

func (s Status) String() string {
	return s.value
}

var (
	StatusScheduled = Status{"scheduled"}
	StatusPublished = Status{"published"}
	StatusDraft     = Status{"draft"}
)

func NewStatusFromString(status string) (Status, error) {
	switch status {
	case StatusPublished.value:
		return StatusPublished, nil
	case StatusScheduled.value:
		return StatusScheduled, nil
	case StatusDraft.value:
		return StatusDraft, nil
	default:
		return Status{}, fmt.Errorf("%s is not a valid status", status)
	}
}

func (s Status) IsZero() bool {
	return s.value == ""
}
