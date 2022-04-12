package logic

import (
	"context"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pb.CallbackRequest) (*pb.CallbackResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CallbackResponse{}, nil
}
