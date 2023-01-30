package common

import (
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

var COSCLIENT *cos.Client

// InitCOSClient 初始化COS
func InitCOSClient() {

	u, _ := url.Parse("https://xxx")
	su, _ := url.Parse("https://xxx")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("xxx"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv("xxx"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	COSCLIENT = client
}

// SaveFile 保存文件
func SaveFile(file *multipart.FileHeader, c *gin.Context) string {
	open, err := file.Open()
	if err != nil {
		log.Printf("SaveFile|文件打开失败|%v", err)
		return ""
	}
	_, err = COSCLIENT.Object.Put(c, "public/"+file.Filename, open, nil)
	if err != nil {
		log.Printf("SaveFile|文件存储失败|%v", err)
		return ""
	}
	returnUrl := COSCLIENT.Object.GetObjectURL("public/" + file.Filename)
	log.Printf("访问url为：%v", returnUrl.String())

	return returnUrl.String()
}
