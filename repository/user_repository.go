package repository

import (
	"gin-rest/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByID(id uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uint) (entity.User, error) {

	var user entity.User
	err := r.db.First(&user, id).Error

	return user, err
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {

	var user entity.User

	err := r.db.Where("email= ?", email).First(&user).Error

	return user, err
}
