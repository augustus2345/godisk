package helper

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go_disk/core/define"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	//id
	// identity
	//name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		fmt.Println("这里也有")
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "cjy <2812485760@qq.com>"
	e.To = []string{mail}
	e.Subject = "验证码测试"
	e.HTML = []byte("验证码是：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "2812485760@qq.com", "ugencsnkcxyadcif", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

// GetCode 生成包含数字和大写字母的4位验证码
func GetCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i <= define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// GetUuid 生成 uuid
func GetUuid() string {
	v4 := uuid.NewV4().String()
	return v4
}

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "go_disk/" + GetUuid() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// MinIOUpload 上传到自建的minio中
//func MinIOUpload(r *http.Request) (string, error) {
//	minioClient, err := minio.New(define.MinIOEndpoint, &minio.Options{
//		Creds: credentials.NewStaticV4(define.MinIOAccessKeyID, define.MinIOAccessSecretKey, ""),
//	})
//	if err != nil {
//		return "", err
//	}
//
//	// 获取文件信息
//	file, fileHeader, err := r.FormFile("file")
//	bucketName := "cloud-disk"
//	objectName := UUID() + path.Ext(fileHeader.Filename)
//
//	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, fileHeader.Size,
//		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
//	if err != nil {
//		return "", err
//	}
//	return define.MinIOBucket + "/" + bucketName + "/" + objectName, nil
//}
