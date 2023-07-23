package login_and_signup

import (
	"github.com/gin-gonic/gin"
	"ini/handler"
	"ini/model/method"
	model "ini/model/user_struct"
)

// @Summary 账号密码登录
// @Description 登录并获得token
// @tags login
// @Accept application/json
// @Param Request body model.LoginU1 true "登录结构体"
// @Success 200 {object} handler.Response "{"message":"登陆成功，获得token"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/login [post]
func LoginWithPwd(c *gin.Context) {
	var LoginUser model.User
	err := c.BindJSON(&LoginUser)
	if err != nil {
		handler.SendBadResponse(c, "登录信息提供出错", err)
		return
	}
	token, err, state := method.UserLogin1(LoginUser)
	switch state {
	case 1:
		{
			handler.SendBadResponse(c, "ID或密码错误", err)
			return
		}
	case 2:
		{
			handler.SendError(c, "创建token出错", err)
			return
		}
	}
	handler.SendGoodResponse(c, "登录成功，获得token", token)
	return
}

// @Summary 发送邮箱验证码
// @Description 向用户提供的邮箱地址发送一个验证码
// @Tags login
// @Accepts application/json
// @Param Request body model.LoginU2 true "邮箱地址"
// @Success 200 {object} handler.Response "{"message":"登陆成功，获得token"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/login [post]
func SendCode(c *gin.Context) {

}

// @Summary 邮箱验证码登录
// @Description 通过邮箱验证码登录，验证码正确则获得token（注：该接口用于发送验证码之后）
// @Tags login
// @Accepts application/json
// @Param Request body model.LoginU2 true "登录结构体"
// @Success 200 {object} handler.Response "{"message":"登陆成功，获得token"}"
// @Failure 400 {object} handler.Response "{"message":"Failure"}"
// @Router /api/v1/login [post]
func LoginWithCode(c *gin.Context) {

}
