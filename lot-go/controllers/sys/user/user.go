package controller

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"meilian/constants"
	controller "meilian/controllers/base"
	userModels "meilian/models/sys/user"

	_ "meilian/models/sys/user"
	"meilian/utils/jwt"
	"unicode/utf8"
)

type UserController struct {
	controller.BaseController
}

func (c *UserController) Login() {
	requestBody := c.Ctx.Input.RequestBody
	print(string(requestBody))
	var loginParams userModels.User
	err := json.Unmarshal(requestBody, &loginParams)
	if err != nil {
		c.ResponseJSON(constants.GainPostDataError, err.Error(), "")
		return
	}
	userName := loginParams.UserName
	password := loginParams.Pwd

	_ = loginParams.Email
	//if flag == "" {
	//	c.ResponseJSON(constants.PARAM_ERROR_CODE,"缺少参数","")
	//}

	if userName == "" {
		c.ResponseJSON(constants.ParamErrorCode, "请输入用户名", "")
		return
	}
	if password == "" {
		c.ResponseJSON(constants.ParamErrorCode, "请输入密码", "")
		return
	}
	//o := orm.NewOrm()
	userModel := userModels.User{UserName: userName}

	err = c.O.Read(&userModel, "UserName")
	if err == orm.ErrNoRows {
		c.ResponseJSON(1, "用户不存在！", "")
		return
	}
	if (userModel.UserName) == "" {
		c.ResponseJSON(constants.ModelNotExistCode, "用户不存在,请先注册！", "")
		return
	}
	if userModel.Pwd != password {
		c.ResponseJSON(constants.PasswordErrorCode, "密码错误，请重试！", "")
		return
	}
	//颁发token令牌
	userInfo := make(map[string]interface{})
	userInfo["id"] = userModel.Id
	userInfo["role"] = userModel.Role
	userInfo["userName"] = userModel.UserName
	token, err := jwt.GenerateToken(userInfo, constants.TOKEN_EXPIRE_TIME)
	if err != nil {
		c.ResponseJSON(constants.TokenGenerateError, "系统异常", "")
		return
	}
	c.ResponseJSON(constants.SuccessCode, "登录成功", token)
}
func (c *UserController) Register() {
	requestBody := c.Ctx.Input.RequestBody
	var registerParams userModels.User
	err := json.Unmarshal(requestBody, &registerParams)
	if err != nil {
		c.ResponseJSON(constants.GainPostDataError, err.Error(), "")
		return
	}

	userName := registerParams.UserName
	password := registerParams.Pwd

	email := registerParams.Email
	userNameLength := utf8.RuneCountInString(userName)
	passwordLength := utf8.RuneCountInString(password)

	if password == "" {
		c.ResponseJSON(constants.ParamErrorCode, "密码不可为空！", "")
		return
	}
	if email == "" {

		c.ResponseJSON(constants.ParamErrorCode, "邮箱不可为空！", "")
		return
	}
	if userNameLength > 25 {
		c.ResponseJSON(constants.ParamErrorCode, "用户名不可以超过25个字！", "")
		return
	}
	if passwordLength > 16 {
		c.ResponseJSON(constants.ParamErrorCode, "请输入16位以内的密码！", "")
		return
	}

	if err != nil {
		c.ResponseJSON(constants.InsertUserError, "系统异常,请联系管理员！", "")
		return
	}
	var existedUserModel userModels.User
	err = c.O.QueryTable(new(userModels.User)).Filter("UserName", userName).One(&existedUserModel)
	if existedUserModel.UserName != "" {
		c.ResponseJSON(constants.ModelHasExistCode, "用户名已存在，换个试试！", "")
		return
	}
	if existedUserModel.Email == email {
		c.ResponseJSON(constants.ModelHasExistCode, "邮箱已存在，换个试试！", "")
		return
	}
	userModel := userModels.User{
		UserName: userName,
		Pwd:      password,
		Email:    email,
	}
	_, err = c.O.Insert(&userModel)
	//颁发token令牌
	userInfo := make(map[string]interface{})
	userInfo["id"] = userModel.Id
	userInfo["role"] = userModel.Role
	userInfo["userName"] = userModel.UserName
	token, err := jwt.GenerateToken(userInfo, constants.TOKEN_EXPIRE_TIME)
	if err != nil {
		c.ResponseJSON(constants.TokenGenerateError, "系统异常", "")
	}

	c.ResponseJSON(1, "注册成功！即将自动登录", token)

}
