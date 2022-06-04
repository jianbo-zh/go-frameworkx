package userrepo

import (
	"goframeworkx/cmd/http/module/usermod/dao"
	"goframeworkx/cmd/http/module/usermod/model"
)

func User(phone string) (user model.User, err error) {
	err = dao.DB().Where("phone=?", phone).First(&user).Error
	return user, err
}
