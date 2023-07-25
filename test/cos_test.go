package test

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go_disk/core/define"
	"net/http"
	"net/url"
	"testing"
)

func TestUpload(t *testing.T) {
	u, _ := url.Parse("https://godisk-1314961495.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: define.TencentSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	key := "go_disk_test/im.png"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/test.png", nil,
	)
	if err != nil {
		panic(err)
	}

}
