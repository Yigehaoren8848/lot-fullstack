package controllers

import (
	"fmt"
	"meilian/constants"
	controller "meilian/controllers/base"
	"meilian/utils/jwt"
	"meilian/utils/upload"
	"os"
	"path/filepath"
	"time"
)

type TestController struct {
	controller.BaseController
}

func (c *TestController) GenerateJwt() {
	data := make(map[string]interface{}, 5)
	data["name"] = "zhang"
	data["id"] = "123"
	data["role"] = "4"
	expiration := 30 * 24 * time.Hour
	jwtString, err := jwt.GenerateToken(data, expiration)
	if err != nil {
		c.ResponseJSON(100, err.Error(), "")
	}
	c.ResponseJSON(1001, "ok", jwtString)
}

func (c *TestController) GetJwtData() {
	tokenString := c.GetString("token")
	if tokenString == "" {
		c.ResponseJSON(0001, "参数异常", "")
	}

	d, err := jwt.GetTokenData(tokenString)

	if err != nil {
		c.ResponseJSON(1, err.Error(), "")
	}
	c.ResponseJSON(1, "ok", d)

}

func (c *TestController) UploadImg() {
	imageFile, header, err := c.GetFile("image")
	orginFileName := header.Filename
	fileExt := filepath.Ext(orginFileName)
	fmt.Print("原始文件名：", fileExt)

	imageFileName := upload.RandomImageFileName() + fileExt
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "文件上传异常,请联系管理员！", "")
		return
	}
	defer imageFile.Close()
	tmpFileName := imageFileName
	c.SaveToFile("image", tmpFileName)
	defer os.Remove(tmpFileName)
	// 读取文件内容
	fileBytes, err := os.ReadFile(tmpFileName)
	if err != nil {
		c.ResponseJSON(constants.ParamErrorCode, "文件上传异常,请联系管理员！", "")
		return
	}
	data, err := upload.UploadImg(imageFileName, fileBytes)
	if err != nil {
		c.ResponseJSON(-1, err.Error(), "")
		return
	}
	//os.ReadFile()
	c.ResponseJSON(1, "ok", data)

}
