package database

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventType(t *testing.T) {

	//arrange
	object, _ := entity.NewEventType("User Login")
	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewEventTypeDAO(db)
	err = dao.Create(object)

	//assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)
	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Description, found.Description)

}

func TestUpdateEventType(t *testing.T) {

	//arrange
	object, _ := entity.NewEventType("User Login")
	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewEventTypeDAO(db)
	err = dao.Create(object)

	object.Description = "User Logoff"
	object.UpdatedAt = time.Now()
	err = dao.Update(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.Nil(t, err)

	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, "User Logoff", found.Description)

}

func TestDeleteEventType(t *testing.T) {

	//arrange
	object, _ := entity.NewEventType("User Login")
	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewEventTypeDAO(db)
	err = dao.Create(object)

	object.Description = "User Logoff"
	err = dao.Delete(object)
	// assert
	assert.NotNil(t, dao)
	assert.Nil(t, err)

	found, err := dao.GetById(object.Id.String())
	assert.NotNil(t, err)

	assert.Nil(t, found)

}

func TestGetByIdEventType(t *testing.T) {

	//arrange
	object, _ := entity.NewEventType("User Login")
	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewEventTypeDAO(db)
	err = dao.Create(object)
	found, err := dao.GetById(object.Id.String())

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, object.Id, found.Id)
	assert.Equal(t, object.Description, found.Description)

}

func TestGetAllEventType(t *testing.T) {

	//arrange
	object, _ := entity.NewEventType("User test")
	db, err := getDBConnection()
	assert.Nil(t, err)

	//act
	dao := NewEventTypeDAO(db)
	err = dao.Create(object)
	foundList, err := dao.GetAll()

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, foundList)
	assert.Greater(t, len(foundList), 0)
	assert.Equal(t, object.Description, foundList[len(foundList)-1].Description)

}

func getDBConnection() (*sql.DB, error) {

	DB_DRIVER := "postgres"
	DB_HOST := "localhost"
	DB_PORT := "5432"
	DB_USER := "root"
	DB_PASSWORD := "root"
	DB_NAME := "userevents"

	//starting database connection
	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println(DataSourceName)

	db, err := sql.Open(DB_DRIVER, DataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
