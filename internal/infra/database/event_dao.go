package database

import (
	"database/sql"

	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
)

type EventDAO struct {
	DB *sql.DB
}

func NewEventDAO(db *sql.DB) *EventDAO {
	return &EventDAO{
		DB: db,
	}
}

func (dao *EventDAO) Create(record *entity.Event) error {

	_, err := dao.DB.Exec("insert into events(id,user_id,event_type_id,performed_at) values($1,$2,$3,$4)", record.Id, record.User.Id, record.EventType.Id, record.PerformedAt)

	return err
}

func (dao *EventDAO) Update(record *entity.Event) error {

	_, err := dao.DB.Exec("update events set user_id=$1,event_type_id=$2,performed_at=$3 where id=$4", record.User.Id, record.EventType.Id, record.PerformedAt, record.Id)

	return err
}

func (dao *EventDAO) Delete(record *entity.Event) error {

	_, err := dao.DB.Exec("delete from events where id=$1", record.Id)

	return err
}

func (dao *EventDAO) GetById(id string) (*entity.Event, error) {

	var ev entity.Event
	err := dao.DB.QueryRow("select id,user_id,event_type_id,performed_at from events where id=$1", id).Scan(&ev.Id, &ev.User.Id, &ev.EventType.Id, &ev.PerformedAt)

	return &ev, err
}

func (dao *EventDAO) GetAll() ([]*entity.Event, error) {

	var evList []*entity.Event
	rows, err := dao.DB.Query("select id,user_id,event_type_id,performed_at from events")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ev entity.Event
		err = rows.Scan(&ev.Id, &ev.User.Id, &ev.EventType.Id, &ev.PerformedAt)
		if err != nil {
			return nil, err
		}
		evList = append(evList, &ev)
	}

	return evList, err
}
