package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := User{Id: userId}
	etId := uuid.New()
	et := EventType{Id: etId}

	//act
	ev, err := NewEvent(user, et, time.Now())

	//assert
	assert.Nil(t, err)
	assert.NotEmpty(t, ev.Id)
	assert.NotEmpty(t, ev.PerformedAt)
	assert.Equal(t, ev.User.Id, userId)
	assert.Equal(t, ev.EventType.Id, etId)

}

func TestWhenUserIsRequired(t *testing.T) {

	//arrange
	user := User{}
	etId := uuid.New()
	et := EventType{Id: etId}

	//act
	_, err := NewEvent(user, et, time.Now())

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrUserIsRequired, err)
}

func TestWhenEventTypeIsRequired(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := User{Id: userId}
	et := EventType{}

	//act
	_, err := NewEvent(user, et, time.Now())

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrEventTypeIsRequired, err)
}

func TestWhenDateIsRequired(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := User{Id: userId}
	etId := uuid.New()
	et := EventType{Id: etId}

	var testTime time.Time
	//act
	_, err := NewEvent(user, et, testTime)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrDateIsRequired, err)
}
