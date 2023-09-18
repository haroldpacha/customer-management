package service

import (
	"log"

	"github.com/haroldpacha/customer-management/dto"
	"github.com/haroldpacha/customer-management/entity"
	"github.com/haroldpacha/customer-management/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Get(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed to map %v: ", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Get(userID string) entity.User {
	return service.userRepository.GetUser(userID)
}
