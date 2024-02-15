package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/raulsilva-tech/UserEventsAPI/internal/dto"
)

// CREATE TABLE user_address(
//     id  varchar not null,
//     user_id varchar not null,
//     cep varchar not null,
//     street varchar,
//     city varchar,
//     state varchar,
//     PRIMARY KEY(id)
// );

type UserAddress struct {
	Id        uuid.UUID `json:"id"`
	User      `json:"user"`
	CEP       string    `json:"cep"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ua *UserAddress) Validate() error {
	if ua.CEP == "" {
		return errors.New("cep is required")
	}
	if ua.User.Id == uuid.Nil {
		return ErrUserIsRequired
	}
	return nil
}

func NewUserAddress(user User, cep string) (*UserAddress, error) {

	ua := UserAddress{
		Id:   uuid.New(),
		User: user,
		CEP:  cep,
	}
	return &ua, ua.Validate()
}

func (ua *UserAddress) GetCEPInfo() error {

	chVIACEP := make(chan dto.ViaCEPDTO)
	chBrasilAPI := make(chan dto.BrasilAPIDTO)

	go ViaCEP(ua.CEP, chVIACEP)
	go BrasilAPI(ua.CEP, chBrasilAPI)

	select {
	case cepInfo := <-chVIACEP:
		ua.City = cepInfo.Localidade
		ua.Street = cepInfo.Logradouro
		fmt.Println("Address information gotten by ViaCEP")
	case cepInfo := <-chBrasilAPI:
		ua.City = cepInfo.City
		ua.Street = cepInfo.Street
		ua.State = cepInfo.State
		fmt.Println("Address information gotten by BrasilAPI")
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout")
		return errors.New("timeout")
	}

	return nil
}

func ViaCEP(cep string, ch chan dto.ViaCEPDTO) {

	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer resp.Body.Close()

	var dto dto.ViaCEPDTO

	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Println("VIACEP: \n", dto)

	ch <- dto
}

func BrasilAPI(cep string, ch chan dto.BrasilAPIDTO) {

	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer resp.Body.Close()

	var dto dto.BrasilAPIDTO

	err = json.NewDecoder(resp.Body).Decode(&dto)
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Println("BRASILAPI: \n", dto)

	ch <- dto

}
