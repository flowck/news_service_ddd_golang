package domain

import (
	"errors"

	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func NewIDFromString(value string) (ID, error) {
	if value == "" {
		return ID{}, errors.New("id is empty")
	}

	id, err := uuid.Parse(value)
	if err != nil {
		return ID{}, errors.New("the id is invalid")
	}

	return ID{value: id}, nil
}

func MustNewIDFromString(value string) ID {
	id, err := NewIDFromString(value)
	if err != nil {
		panic(err)
	}

	return id
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) IsZero() bool {
	return i.value.ID() == 0
}
