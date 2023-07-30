package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var COSCLIENT *cos.Client

// InitCOSClient 初始化COS
func InitCOSClient() {

	u, _ := url.Parse("xxx")
	su, _ := url.Parse("xxx")
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

	// 获取当前日期
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	milli := currentTime.UnixMilli()

	// 创建文件夹路径
	folderPath := filepath.Join("mini_douyin", fmt.Sprintf("%d", year), fmt.Sprintf("%02d", month), fmt.Sprintf("%02d", day), "douyin_"+strconv.FormatInt(milli, 10)+"_"+file.Filename)
	folderPath = filepath.ToSlash(folderPath)
	_, err = COSCLIENT.Object.Put(c, folderPath, open, nil)
	if err != nil {
		log.Printf("SaveFile|文件存储失败|%v", err)
		return ""
	}
	returnUrl := COSCLIENT.Object.GetObjectURL(folderPath)
	log.Printf("访问url为：%v", returnUrl.String())

	return returnUrl.String()
}
