package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go_disk/core/helper"
	"go_disk/core/modles"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_disk/core/internal/logic"
	"go_disk/core/internal/svc"
	"go_disk/core/internal/types"
)

func FileUpLoadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 判断当前用户容量是否达到上限
		userIdentity := r.Header.Get("UserIdentity")
		ub := new(modles.UserBasic)
		has, err := svcCtx.Engine.Where("identity = ?", userIdentity).Select("now_volume,total_volume").Get(ub)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if fileHeader.Size+ub.NowVolume > ub.TotalVolume {
			httpx.Error(w, errors.New("已超出当前容量"))
			return
		}
		// 如果没有超出，就判断当前文件是否已经存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(modles.RepositoryPool)
		has, err = svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		// 判断使用的存储引擎，默认cos
		var filePath string
		filePath, err = helper.CosUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 往 logic 中传递 request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = filePath

		l := logic.NewFileUpLoadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpLoad(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
