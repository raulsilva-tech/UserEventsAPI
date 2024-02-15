package database

import (
	"database/sql"

	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
)

type UserAddressDAO struct {
	DB *sql.DB
}

func NewUserAddressDAO(db *sql.DB) *UserAddressDAO {
	return &UserAddressDAO{
		DB: db,
	}
}

func (dao *UserAddressDAO) Create(record *entity.UserAddress) error {

	record.GetCEPInfo()
	_, err := dao.DB.Exec("insert into user_address (id,user_id,cep,street,city,state) values($1,$2,$3,$4,$5,$6)", record.Id, record.User.Id, record.CEP, record.Street, record.City, record.State)

	return err
}

func (dao *UserAddressDAO) Update(record *entity.UserAddress) error {

	record.GetCEPInfo()
	_, err := dao.DB.Exec("update user_address set user_id=$1,cep=$2,street=$3,city=$4,state=$5 where id=$6 ", record.User.Id, record.CEP, record.Street, record.City, record.State, record.Id)

	return err
}

func (dao *UserAddressDAO) Delete(record *entity.UserAddress) error {

	_, err := dao.DB.Exec("delete from user_address where id=$1", record.Id)

	return err
}

func (dao *UserAddressDAO) GetById(id string) (*entity.UserAddress, error) {

	var ua entity.UserAddress
	err := dao.DB.QueryRow("select id,user_id,cep,street,city,state from user_address where id=$1", id).Scan(&ua.Id, &ua.User.Id, &ua.CEP, &ua.Street, &ua.City, &ua.State)

	return &ua, err
}

func (dao *UserAddressDAO) GetAll() ([]*entity.UserAddress, error) {

	var uaList []*entity.UserAddress
	rows, err := dao.DB.Query("select id,user_id,cep,street,city,state from user_address")

	for rows.Next() {

		var ua entity.UserAddress
		err = rows.Scan(&ua.Id, &ua.User.Id, &ua.CEP, &ua.Street, &ua.City, &ua.State)
		if err != nil {
			return nil, err
		}
		uaList = append(uaList, &ua)

	}

	return uaList, err
}
