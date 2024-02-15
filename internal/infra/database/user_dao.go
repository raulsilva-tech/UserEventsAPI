package database

import (
	"database/sql"

	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
)

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{
		DB: db,
	}
}

func (dao *UserDAO) Create(record *entity.User) error {

	_, err := dao.DB.Exec("insert into users(id,name,email,password,created_at,updated_at) values ($1,$2,$3,$4,$5,$6)", record.Id, record.Name, record.Email, record.Password, record.CreatedAt, record.UpdatedAt)

	return err

}

func (dao *UserDAO) Update(record *entity.User) error {

	_, err := dao.DB.Exec("update users set name=$1,email=$2, password=$3,updated_at=$4 where id=$5", record.Name, record.Email, record.Password, record.UpdatedAt, record.Id)

	return err

}

func (dao *UserDAO) Delete(record *entity.User) error {

	_, err := dao.DB.Exec("delete from users where id=$1", record.Id)

	return err

}

func (dao *UserDAO) GetById(id string) (*entity.User, error) {

	var record entity.User

	err := dao.DB.QueryRow("select id,name,email,password,created_at,updated_at from users where id=$1", id).Scan(&record.Id, &record.Name, &record.Email, &record.Password, &record.CreatedAt, &record.UpdatedAt)

	// stmt, err := dao.DB.Prepare("select id,name,email,password,created_at,updated_at from users where id=$1")
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// err = stmt.QueryRow(id).Scan(&record.Id, &record.Name, &record.Email, &record.Password, &record.CreatedAt, &record.UpdatedAt)

	return &record, err
}

func (dao *UserDAO) GetAll() ([]*entity.User, error) {

	var recordList []*entity.User

	rows, err := dao.DB.Query("select id,name,email,password,created_at,updated_at from users")

	for rows.Next() {

		var record entity.User
		err = rows.Scan(&record.Id, &record.Name, &record.Email, &record.Password, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, err
		}
		recordList = append(recordList, &record)
	}

	return recordList, err
}
