package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_disk/core/modles"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/go_disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败。。。")
		t.Fatal(err)
	}
	data := make([]*modles.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		fmt.Println("数据库find失败。。。")
		t.Fatal(err)
	}
	fmt.Println(data)
	b, err := json.Marshal(data)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
