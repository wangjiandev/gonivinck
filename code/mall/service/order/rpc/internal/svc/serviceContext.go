package svc

import (
	"mall/service/order/model"
	"mall/service/order/rpc/internal/config"
	"mall/service/product/rpc/product"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
