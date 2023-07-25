package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "go_disk_key"

// 定义验证码的长度
var CodeLength = 6

// 定义验证码的过期时间
var CodeExpire = 300

// ObjectStorageType 对象存储类型
// 支持 minio\cos
var ObjectStorageType = os.Getenv("ObjectStorageType")

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = os.Getenv("TencentSecretKey")
var TencentSecretID = os.Getenv("TencentSecretID")
var CosBucket = "https://godisk-1314961495.cos.ap-nanjing.myqcloud.com/"
