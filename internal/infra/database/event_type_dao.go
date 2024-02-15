package database

import (
	"database/sql"

	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
)

type EventTypeDAO struct {
	DB *sql.DB
}

func NewEventTypeDAO(db *sql.DB) *EventTypeDAO {
	return &EventTypeDAO{
		DB: db,
	}
}

func (dao *EventTypeDAO) Create(record *entity.EventType) error {

	stmt, err := dao.DB.Prepare("insert into event_types(id,description,updated_at,created_at) values($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(record.Id, record.Description, record.UpdatedAt, record.CreatedAt)
	return err
}

func (dao *EventTypeDAO) Update(record *entity.EventType) error {

	stmt, err := dao.DB.Prepare("update event_types set description = $1, updated_at=$2, created_at=$3 where id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(record.Description, record.UpdatedAt, record.CreatedAt, record.Id)
	return err
}

func (dao *EventTypeDAO) Delete(record *entity.EventType) error {

	stmt, err := dao.DB.Prepare("delete from event_types where id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(record.Id)
	return err
}

func (dao *EventTypeDAO) GetById(id string) (*entity.EventType, error) {

	stmt, err := dao.DB.Prepare("select id,description,updated_at,created_at from event_types where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var et entity.EventType

	err = stmt.QueryRow(id).Scan(&et.Id, &et.Description, &et.UpdatedAt, &et.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &et, nil

}

func (dao *EventTypeDAO) GetAll() ([]*entity.EventType, error) {

	rows, err := dao.DB.Query("select id,description,updated_at,created_at from event_types")
	if err != nil {
		return nil, err
	}

	var etList []*entity.EventType

	for rows.Next() {

		var et entity.EventType
		err = rows.Scan(&et.Id, &et.Description, &et.UpdatedAt, &et.CreatedAt)
		if err != nil {
			return nil, err
		}
		etList = append(etList, &et)

	}

	return etList, nil

}
