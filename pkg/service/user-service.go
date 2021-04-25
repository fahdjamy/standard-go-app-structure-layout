package service

import "github.com/fahdjamy/standard-structure-layout/pkg/types"

type userService struct {}

func (u userService) DeleteUser(uid string) error {
	panic("implement me")
}

func (u userService) FindAll() ([]types.User, error) {
	panic("implement me")
}

func (u userService) Validate(user *types.User) error {
	panic("implement me")
}

func (u userService) Create(user *types.User) (*types.User, error) {
	panic("implement me")
}

func NewUserService() types.UserService {
	return &userService{}
}
