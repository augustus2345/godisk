package logic

import (
	"context"
	"errors"
	"go_disk/core/define"
	"go_disk/core/helper"
	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"
	"go_disk/core/modles"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRegisterRequest) (resp *types.MailCodeSendRegisterReply, err error) {
	// 1. 判断邮箱是否已经被注册
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(modles.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	// 2. 如果没有被注册，就获取验证码
	code := helper.GetCode()
	// 3. 先将验证码存储在 redis 里面，并设置 5 分钟的过期时间
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	// 4. 然后给用户发送验证码
	err = helper.MailSendCode(req.Email, code)
	return
}
