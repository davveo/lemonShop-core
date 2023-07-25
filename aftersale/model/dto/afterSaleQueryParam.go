package dto

type AfterSaleQueryParam struct {
	PageNo        int    `json:"page_no" comment:"页码"`
	PageSize      int    `json:"page_size" comment:"分页大小"`
	MemberID      int    `json:"member_id" comment:"会员ID"`
	SellerID      int    `json:"seller_id" comment:"商家ID"`
	Keyword       string `json:"keyword" comment:"模糊查询的关键字"`
	ServiceSn     string `json:"service_sn" comment:"售后服务单号"`
	OrderSn       string `json:"order_sn" comment:"订单编号"`
	GoodsName     string `json:"goods_name" comment:"商品名称"`
	ServiceType   string `json:"service_type" comment:"售后类型 RETURN_GOODS：退货，CHANGE_GOODS：换货，SUPPLY_AGAIN_GOODS：补发货品，ORDER_CANCEL：取消订单"`
	ServiceStatus string `json:"service_status" comment:"售后单状态 APPLY：申请，PASS：审核通过，REFUSE：审核拒绝，WAIT_FOR_MANUAL：待人工处理，STOCK_IN：入库，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成"`
	StartTime     int64  `json:"start_time" comment:"申请时间-起始时间"`
	EndTime       int64  `json:"end_time" comment:"申请时间-结束时间"`
	CreateChannel string `json:"create_channel" comment:"售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建"`
}
