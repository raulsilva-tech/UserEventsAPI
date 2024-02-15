package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// CREATE table event_types(
//     id  varchar not null,
//     description varchar not null,
//     created_at		timestamp,
// 	updated_at		timestamp,
// 	PRIMARY KEY ( id )
// );

var ErrDescriptionIsRequired = errors.New("description is required")

type EventType struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (et *EventType) Validate() error {
	if et.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}

func NewEventType(description string) (*EventType, error) {
	et := EventType{
		Id:          uuid.New(),
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &et, et.Validate()
}
