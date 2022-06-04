package userapi

import (
	"goframeworkx/cmd/http/module/usermod/api/apicmn"
	"goframeworkx/cmd/http/module/usermod/services/usersvc"

	"github.com/gin-gonic/gin"
	"github.com/jianbo-zh/go-errors"
	goi18n "github.com/jianbo-zh/go-i18n"
)

// @Tags 登录
// @Title 登录
// @Description 登录
// @Param   body     	body    LoginParam  true      "离线消息"
// @Success 200 {object} model.User 操作结果
// @router /cmsprovider/v1/message/uploadOffline [post]
func Login(ctx *gin.Context) {
	var param LoginParam
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		apicmn.Error(ctx, errors.Newc(apicmn.ErrParamErr, goi18n.Sprintf("params error")).With(errors.Inner(err)))
		return
	}

	user, err := usersvc.User(usersvc.UserParam{Phone: param.Phone})
	if err != nil {
		apicmn.Error(ctx, errors.Newc(apicmn.ErrParamErr, goi18n.Sprintf("params error")).With(errors.Inner(err)))
	}

	apicmn.Success(ctx, user)
}
