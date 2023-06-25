package repository

import "iris-boilerplate/entity"

type UserRepository interface {
	FindAll() (user []entity.User)
}
