package service

import (
	"demo/model"
	"github.com/go-xorm/xorm"
)

type UserService interface {
	GetUserInfo (userId int64) model.User
}

func NewUserService(engine *xorm.Engine) UserService {
	return &userService{
		Engine: engine,
	}
}

type userService struct {
	Engine *xorm.Engine
}

func (uc *userService) GetUserInfo(userId int64) model.User {
	var user model.User
	err := uc.Engine.Where("id = %d", userId).Find(&user)
	if err != nil {
		panic(err.Error())
	}
	return user
}
