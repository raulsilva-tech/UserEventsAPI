package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventType(t *testing.T) {

	//arrange act
	et, err := NewEventType("Description")

	//assure
	assert.Nil(t, err)
	assert.NotNil(t, et)
	assert.NotEmpty(t, et.Id)
	assert.Equal(t, et.Description, "Description")
	assert.NotEmpty(t, et.CreatedAt)
	assert.NotEmpty(t, et.UpdatedAt)
}

func TestWhenErrDescriptionIsRequired(t *testing.T) {

	//arrange act
	_, err := NewEventType("")

	//assure
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
