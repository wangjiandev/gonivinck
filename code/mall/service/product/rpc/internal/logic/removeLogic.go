package logic

import (
	"context"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	err := l.svcCtx.ProductModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &pb.RemoveResponse{}, nil
}
