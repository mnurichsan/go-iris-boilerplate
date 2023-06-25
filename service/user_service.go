package service

import "iris-boilerplate/model"

type UserService interface {
	List() (responses []model.GetUserResponse)
}
