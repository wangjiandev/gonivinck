package logic

import (
	"context"

	"mall/service/product/model"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/pb"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateRequest) (*pb.CreateResponse, error) {
	newProduct := model.Product{}
	copier.Copy(&newProduct, in)
	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newProduct.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &pb.CreateResponse{
		Id: newProduct.Id,
	}, nil
}
