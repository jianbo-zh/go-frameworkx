package usersvc

import (
	"goframeworkx/cmd/http/module/usermod/model"
	"goframeworkx/cmd/http/module/usermod/repo/userrepo"
)

func User(user UserParam) (model.User, error) {
	return userrepo.User(user.Phone)
}
