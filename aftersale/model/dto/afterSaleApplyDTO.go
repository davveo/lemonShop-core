package dto

import (
	"github.com/davveo/lemonShop-core/trade/order/model"
	"github.com/davveo/lemonShop-framework/security/entity"
)

type AfterSaleApplyDTO struct {
	Buyer      entity.Buyer      // 当前登录的会员信息
	Order      models.Order      // 申请售后的订单信息
	OrderItems models.OrderItems // 申请售后的订单商品信息
}
