package dto

type ServiceOperateAllowable struct {
	AllowAudit             bool `json:"allow_audit"`               // 是否允许商家审核
	AllowShip              bool `json:"allow_ship"`                // 是否允许用户退还商品（填充物流信息）
	AllowPutInStore        bool `json:"allow_put_in_store"`        // 是否允许商家入库
	AllowAdminRefund       bool `json:"allow_admin_refund"`        // 是否允许平台退款
	AllowShowStorageNum    bool `json:"allow_show_storage_num"`    // 是否允许显示入库数量
	AllowShowReturnAddr    bool `json:"allow_show_return_addr"`    // 是否允许显示退货地址
	AllowShowShipInfo      bool `json:"allow_show_ship_info"`      // 是否允许显示用户填写的物流信息
	AllowSellerRefund      bool `json:"allow_seller_refund"`       // 是否允许商家退款
	AllowSellerCreateOrder bool `json:"allow_seller_create_order"` // 是否允许商家手动创建新订单
	AllowSellerClose       bool `json:"allow_seller_close"`        // 是否允许商家关闭售后服务
}
