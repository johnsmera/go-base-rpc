package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/johnsmera/chall/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during the user validation %v", err)
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
	}

	return user, err
}
