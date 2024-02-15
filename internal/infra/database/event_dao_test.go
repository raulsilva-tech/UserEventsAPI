package database

import (
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {

	//arrange
	db, err := getDBConnection()
	assert.Nil(t, err)

	userDAO := NewUserDAO(db)
	etDAO := NewEventTypeDAO(db)

	user := entity.User{Id: uuid.New(), Email: "teste1@mail.com"}
	et := entity.EventType{Id: uuid.New()}
	object, err := entity.NewEvent(user, et, time.Now())

	//act
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	err = etDAO.Create(&et)
	assert.Nil(t, err)

	dao := NewEventDAO(db)
	err = dao.Create(object)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.User.Id, found.User.Id)
	assert.Equal(t, object.EventType.Id, found.EventType.Id)

}

func TestUpdateEvent(t *testing.T) {

	//arrange
	db, err := getDBConnection()
	assert.Nil(t, err)

	userDAO := NewUserDAO(db)
	etDAO := NewEventTypeDAO(db)

	user := entity.User{Id: uuid.New(), Email: "teste22@mail.com"}
	et := entity.EventType{Id: uuid.New()}
	firstTime := time.Now()
	object, err := entity.NewEvent(user, et, firstTime)

	//act
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	err = etDAO.Create(&et)
	assert.Nil(t, err)

	dao := NewEventDAO(db)
	err = dao.Create(object)

	newType := &entity.EventType{Id: uuid.New()}
	err = etDAO.Create(newType)

	object.EventType.Id = newType.Id
	err = dao.Update(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, newType.Id.String(), found.EventType.Id.String())

}

func TestDeleteEvent(t *testing.T) {

	//arrange
	db, err := getDBConnection()
	assert.Nil(t, err)

	userDAO := NewUserDAO(db)
	etDAO := NewEventTypeDAO(db)

	user := entity.User{Id: uuid.New(), Email: "teste4@mail.com"}
	et := entity.EventType{Id: uuid.New()}
	object, err := entity.NewEvent(user, et, time.Now())

	//act
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	err = etDAO.Create(&et)
	assert.Nil(t, err)

	dao := NewEventDAO(db)
	err = dao.Create(object)
	assert.Nil(t, err)
	err = dao.Delete(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.NotNil(t, err)

	assert.Equal(t, found.Id, uuid.Nil)

}

func TestGetByIdEvent(t *testing.T) {

	//arrange
	db, err := getDBConnection()
	assert.Nil(t, err)

	userDAO := NewUserDAO(db)
	etDAO := NewEventTypeDAO(db)

	user := entity.User{Id: uuid.New(), Email: "teste5@mail.com"}
	et := entity.EventType{Id: uuid.New()}
	object, err := entity.NewEvent(user, et, time.Now())

	//act
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	err = etDAO.Create(&et)
	assert.Nil(t, err)

	dao := NewEventDAO(db)
	err = dao.Create(object)
	found, err := dao.GetById(object.Id.String())

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, object.Id, found.Id)

}

func TestGetAllEvent(t *testing.T) {

	//arrange
	db, err := getDBConnection()
	assert.Nil(t, err)

	userDAO := NewUserDAO(db)
	etDAO := NewEventTypeDAO(db)

	user := entity.User{Id: uuid.New(), Email: "teste100@mail.com"}
	et := entity.EventType{Id: uuid.New()}
	object, err := entity.NewEvent(user, et, time.Now())

	//act
	err = userDAO.Create(&user)
	assert.Nil(t, err)
	err = etDAO.Create(&et)
	assert.Nil(t, err)

	dao := NewEventDAO(db)
	err = dao.Create(object)

	foundList, err := dao.GetAll()

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, foundList)
	assert.Greater(t, len(foundList), 0)

}
