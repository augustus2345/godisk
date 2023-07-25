package logic

import (
	"context"
	"fmt"
	"go_disk/core/modles"

	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest) (resp *types.UserRepositorySaveReply, err error) {
	// 判断文件是否超容量
	rp := new(modles.RepositoryPool)
	_, err := l.svcCtx.Engine.Select("size").Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		fmt.Println(err)
		return
	}
	ub := new(modles.UserBasic)
	_, err = l.svcCtx.Engine.Select("now_volum,total_valum").Where("identity=?", us)
	return
}
