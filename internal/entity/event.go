package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// CREATE TABLE events(
// 	id				varchar NOT NULL,
// 	user_id			varchar NOT NULL,
//   event_type_id   varchar not null,
//   performed_at	timestamp,
//   PRIMARY KEY ( id )
// );

var (
	ErrUserIsRequired      = errors.New("user is required")
	ErrEventTypeIsRequired = errors.New("event type is required")
	ErrDateIsRequired      = errors.New("'performed at' is required")
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	User        `json:"user"`
	EventType   `json:"event_type"`
	PerformedAt time.Time `json:"performed_at"`
}

func (e *Event) Validate() error {
	if e.User.Id == uuid.Nil {
		return ErrUserIsRequired
	}
	if e.EventType.Id == uuid.Nil {
		return ErrEventTypeIsRequired
	}
	if e.PerformedAt.IsZero() {
		return ErrDateIsRequired
	}
	return nil
}

func NewEvent(user User, eventType EventType, performedAt time.Time) (*Event, error) {

	e := Event{
		Id:          uuid.New(),
		User:        User{Id: user.Id},
		EventType:   EventType{Id: eventType.Id},
		PerformedAt: performedAt,
	}

	return &e, e.Validate()
}
