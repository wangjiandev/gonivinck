package logic

import (
	"context"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *pb.DetailRequest) (*pb.DetailResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DetailResponse{}, nil
}
