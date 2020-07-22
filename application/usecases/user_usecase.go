package usecases

import (
	"log"

	"github.com/leopedroso45/codeedu/application/repositories"
	"github.com/leopedroso45/codeedu/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Insert(user)
	if err != nil {
		log.Fatalf("Error during userRepository creating: %v", err)
		return user, err
	}
	return user, nil
}
