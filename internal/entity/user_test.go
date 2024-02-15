package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	//arrange and act
	user, err := NewUser("User", "user@gmail.com", "123", time.Now(), time.Now())

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
	assert.NotEqual(t, "123", user.Password)
}

func TestWhenNameIsRequired(t *testing.T) {

	_, err := NewUser("", "user@gmail.com", "123", time.Now(), time.Now())

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)

}

func TestWhenEmailIsRequired(t *testing.T) {

	_, err := NewUser("User", "", "123", time.Now(), time.Now())

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrEmailIsRequired, err)

}

func TestWhenPasswordIsRequired(t *testing.T) {

	_, err := NewUser("User", "user@gmail.com", "", time.Now(), time.Now())

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrPasswordIsRequired, err)

}
