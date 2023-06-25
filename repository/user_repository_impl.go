package repository

import (
	"iris-boilerplate/entity"

	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *userRepositoryImpl) FindAll() (users []entity.User) {
	err := repository.DB.Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users
}
