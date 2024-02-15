package dto

import "time"

type CreateEventTypeInput struct {
	Description string `json:"description"`
}

type CreateEventInput struct {
	UserId      string    `json:"user_id"`
	EventTypeId string    `json:"event_type_id"`
	PerformedAt time.Time `json:"performed_at"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserAddressInput struct {
	UserId string `json:"user_id"`
	CEP    string `json:"cep"`
}

type ViaCEPDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPIDTO struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}
