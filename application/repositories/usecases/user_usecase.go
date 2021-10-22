package usecases

import (
	"log"

	"github.com/johnsmera/chall/application/repositories"
	"github.com/johnsmera/chall/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (strucd *UserUseCase) Create(user *domain.User) (*domain.User, error) {

	user, err := strucd.UserRepository.Insert(user)

	if err != nil {
		log.Fatalf("Error during create user at repository %v", err)
		return user, err
	}

	return user, nil
}
