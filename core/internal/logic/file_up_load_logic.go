package logic

import (
	"context"

	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUpLoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUpLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUpLoadLogic {
	return &FileUpLoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUpLoadLogic) FileUpLoad(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	// todo: add your logic here and delete this line

	return
}
