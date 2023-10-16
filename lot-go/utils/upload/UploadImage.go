package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"meilian/constants"
	"strconv"
	"time"
)

func UploadImg(fileName string, fileBytes []byte) (imgUrl string, err error) {
	qiniuAccessKey := constants.QiniuAccessKey
	qiniuSecretKey := constants.QiniuSecretKey
	qiniuBuketName := constants.QiniuBuketName
	//七牛云凭证
	mac := qbox.NewMac(qiniuAccessKey, qiniuSecretKey)
	putPolicy := storage.PutPolicy{
		Scope: qiniuBuketName,
	}
	uptoken := putPolicy.UploadToken(mac)
	//localFile := "E:/project/vue/ruang-admin/img/indoapr20.jpg"

	// 初始化配置
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err = formUploader.Put(context.Background(), &ret, uptoken, fileName, bytes.NewReader(fileBytes), int64(len(fileBytes)), nil)
	//err = formUploader.PutFile(context.Background(), &ret, uptoken, fileName, bytes.NewReader(fileBytes), int64(len(fileBytes)),nil)
	if err != nil {
		fmt.Println("上传图片失败:", err)
		return
	}
	// 输出图片外链地址
	imageURL := ret.Key
	fmt.Println("图片上传成功，外链地址：", imageURL)
	domain := constants.QiniuImageDomain
	//publicAccessURL := storage.MakePublicURL(domain, imageURL)
	deadline := time.Now().Add(time.Second * 3600).Unix()

	privateAccessURL := storage.MakePrivateURL(mac, domain, fileName, deadline)
	fmt.Println(privateAccessURL)

	return imageURL, err
}

func RandomImageFileName() string {
	// 获取当前时间的时间戳，单位是纳秒
	currentTime := time.Now().UnixNano()

	// 将时间戳转换为秒，可以使用以下代码
	// currentTime := time.Now().Unix()

	filename := strconv.FormatInt(currentTime, 10)
	return string(filename)
}
