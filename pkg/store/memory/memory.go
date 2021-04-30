package memory

import (
	"errors"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
)

type Database struct {
	Users []types.DBData
}

func (d *Database) Delete(propertyId string) error {
	return nil
}

func (d *Database) GetAll(model string) ([]types.DBData, error) {
	switch model {
	case "user":
		return d.Users, nil
	default:
		return nil, errors.New("model name does not exist")
	}
}

func (d *Database) FindById(model string, propertyId string) (types.DBData, error) {
	var response types.DBData
	switch model {
	case "user":
		for _, user := range d.Users{
			u := user.(types.User)
			if u.ID == propertyId {
				response = u
				return response, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (d *Database) Create(model string, data *types.DBData) (types.DBData, error) {
	var response types.DBData
	switch model {
	case "user":
		_ = append(d.Users, *data)
	default:
		return nil, errors.New("model name does not exist")
	}
	return response, nil
}

func NewInMemoryDB() types.DBService {
	return &Database{}
}
