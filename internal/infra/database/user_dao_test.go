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

func TestCreateUser(t *testing.T) {

	//arrange
	object, _ := entity.NewUser("User", "user10@gmail.com", "123", time.Now(), time.Now())

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewUserDAO(db)
	err = dao.Create(object)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Name, found.Name)

}

func TestUpdateUser(t *testing.T) {

	//arrange
	object, _ := entity.NewUser("User", "user12@gmail.com", "123", time.Now(), time.Now())

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewUserDAO(db)
	err = dao.Create(object)

	object.Name = "User Logoff"
	object.UpdatedAt = time.Now()
	err = dao.Update(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, "User Logoff", found.Name)

}

func TestDeleteUser(t *testing.T) {

	//arrange
	object, _ := entity.NewUser("User", "user13@gmail.com", "123", time.Now(), time.Now())

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewUserDAO(db)
	err = dao.Create(object)

	object.Name = "User Logoff"
	err = dao.Delete(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.NotNil(t, err)

	assert.Equal(t, uuid.Nil, found.Id)

}

func TestGetByIdUser(t *testing.T) {

	//arrange
	object, _ := entity.NewUser("User", "user3@gmail.com", "123", time.Now(), time.Now())

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewUserDAO(db)
	err = dao.Create(object)
	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Name, found.Name)

}

func TestGetAllUser(t *testing.T) {

	//arrange
	object, _ := entity.NewUser("User", "user@gmail.com", "123", time.Now(), time.Now())

	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewUserDAO(db)
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
