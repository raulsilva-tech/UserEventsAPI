package graph

import "github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserDAO        *database.UserDAO
	UserAddressDAO *database.UserAddressDAO
	EventTypeDAO   *database.EventTypeDAO
	EventDAO       *database.EventDAO
}
