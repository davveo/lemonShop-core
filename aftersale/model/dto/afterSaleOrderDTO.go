package dto

import models "github.com/davveo/lemonShop-core/promotion/fulldiscount/model"

type AfterSaleOrderDTO struct {
	OrderSn          string                    `json:"order_sn"`           // 订单编号
	GoodID           int                       `json:"good_id"`            // 商品id
	SkuID            int                       `json:"sku_id"`             // 产品id
	GoodsName        string                    `json:"goods_name"`         // 商品名称
	GoodsImg         string                    `json:"goods_img"`          // 商品图片
	GoodsPrice       float64                   `json:"goods_price"`        // 商品单价
	BuyNum           int                       `json:"buy_num"`            // 购买数量
	ProvinceID       int                       `json:"province_id"`        // 收货地址省份ID
	CityID           int                       `json:"city_id"`            // 收货地址城市ID
	CountyID         int                       `json:"county_id"`          // 收货地址县(区)ID
	TownID           int                       `json:"town_id"`            // 收货地址乡(镇)ID
	Province         string                    `json:"province"`           // 收货地址省份名称
	City             string                    `json:"city"`               // 收货地址城市名称
	County           string                    `json:"county"`             // 收货地址县(区)名称
	Town             string                    `json:"town"`               // 收货地址城镇名称
	ShipAddr         string                    `json:"ship_addr"`          // 收货地址详细
	ShipName         string                    `json:"ship_name"`          // 收货人姓名
	ShipMobile       string                    `json:"ship_mobile"`        // 收货人手机号
	GiftList         []models.FullDiscountGift `json:"gift_list"`          // 订单赠品数据集合
	IsRetrace        bool                      `json:"is_retrace"`         // 是否支持原路退回
	IsReceipt        bool                      `json:"is_receipt"`         // 是否有发票
	AllowReturnGoods bool                      `json:"allow_return_goods"` // 是否允许申请退货
	SellerName       string                    `json:"seller_name"`        // 商家名称
}
