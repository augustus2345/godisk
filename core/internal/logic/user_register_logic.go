package logic

import (
	"context"
	"errors"
	"go_disk/core/helper"
	"go_disk/core/modles"
	"log"

	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// 1. 注册时，验证码是否正确
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取改邮箱验证码")

	}
	if code != req.Code {
		// 验证码不正确
		err = errors.New("验证码错误")
		return
	}
	// 2. 判断注册时，用户名是否重复
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(modles.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("用户名已存在")
		return
	}
	// 3. 数据存入数据库
	user := &modles.UserBasic{
		Identity: helper.GetUuid(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println(n)
	return
}
