package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := entity.User{Id: userId, Email: "test@mail.com"}

	object, err := entity.NewUserAddress(user, "11704140")

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	userDAO := NewUserDAO(db)
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	dao := NewUserAddressDAO(db)
	err = dao.Create(object)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Name, found.Name)

}

func TestUpdateUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := entity.User{Id: userId, Email: "testu@mail.com"}

	object, err := entity.NewUserAddress(user, "11704140")

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	userDAO := NewUserDAO(db)
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	dao := NewUserAddressDAO(db)
	err = dao.Create(object)

	object.CEP = "01560020"
	object.UpdatedAt = time.Now()
	err = dao.Update(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.CEP, found.CEP)

}

func TestDeleteUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := entity.User{Id: userId, Email: "testd@mail.com"}

	object, err := entity.NewUserAddress(user, "11704140")

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	userDAO := NewUserDAO(db)
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	dao := NewUserAddressDAO(db)
	err = dao.Create(object)

	err = dao.Delete(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.NotNil(t, err)

	assert.Equal(t, uuid.Nil, found.Id)

}

func TestGetByIdUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := entity.User{Id: userId, Email: "testg@mail.com"}

	object, err := entity.NewUserAddress(user, "11704140")

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	userDAO := NewUserDAO(db)
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	dao := NewUserAddressDAO(db)
	err = dao.Create(object)

	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Name, found.Name)

}

func TestGetAllUserAddress(t *testing.T) {

	//arrange
	userId := uuid.New()
	user := entity.User{Id: userId, Email: "testall@mail.com"}

	object, err := entity.NewUserAddress(user, "11704140")

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	userDAO := NewUserDAO(db)
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	dao := NewUserAddressDAO(db)
	err = dao.Create(object)

	foundList, err := dao.GetAll()

	for _, found := range foundList {
		fmt.Println(found)
	}

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, foundList)
	assert.Greater(t, len(foundList), 0)
	assert.Equal(t, object.Name, foundList[len(foundList)-1].Name)

}
