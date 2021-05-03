package types

import (
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Router interface {
	Serve (port string)
	NotFound (f func(w http.ResponseWriter, r *http.Request))
	Get (uri string, f func(w http.ResponseWriter, r *http.Request))
	Post (uri string, f func(w http.ResponseWriter, r *http.Request))
}

type DBData interface {
	ValidateData() error
}

type DBService interface {
	Delete(propertyId string) error
	GetAll(model string) ([]DBData, error)
	Create(model string, data *DBData) (DBData, error)
	FindById(model string, propertyId string) (DBData, error)
}
