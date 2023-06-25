package service

import (
	"iris-boilerplate/model"
	"iris-boilerplate/repository"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) List() (responses []model.GetUserResponse) {
	users := service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return responses
}
