package entity

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWhenGetCEPInfo(t *testing.T) {
	// arrange
	userId := uuid.New()
	user := User{Id: userId}

	// act
	ua, _ := NewUserAddress(user, "11704140")

	err := ua.GetCEPInfo()
	fmt.Printf("City: %s , Street: %s", ua.City, ua.Street)
	// assert
	assert.Nil(t, err)
	assert.NotNil(t, ua)
	assert.NotEmpty(t, ua.City)
	assert.NotEmpty(t, ua.Street)
}

func TestNewUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := User{Id: userId}

	//act
	ua, err := NewUserAddress(user, "11704140")

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, ua)
	assert.NotEmpty(t, ua.Id)
	assert.NotEmpty(t, ua.User.Id)
	assert.NotEmpty(t, ua.CEP)

}

func TestWhenUserFromUserAddressIsRequired(t *testing.T) {

	//arrange
	userId := uuid.Nil
	user := User{Id: userId}

	//act
	_, err := NewUserAddress(user, "11704140")

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, ErrUserIsRequired, err)

}
