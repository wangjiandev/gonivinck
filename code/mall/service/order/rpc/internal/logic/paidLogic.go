package logic

import (
	"context"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *pb.PaidRequest) (*pb.PaidResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.PaidResponse{}, nil
}
