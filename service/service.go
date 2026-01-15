package service

import "gin-rest/entity"

type UserService interface {
	Create(user entity.User) (entity.User, error)
	UserFindByID(id uint) (entity.User, error)
}
