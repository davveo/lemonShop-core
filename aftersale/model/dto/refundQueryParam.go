package dto

type RefundQueryParam struct {
	PageNo        int    `json:"page_no"`        // 页码
	PageSize      int    `json:"page_size"`      // 分页大小
	MemberId      int    `json:"member_id"`      // 会员ID
	SellerId      int    `json:"seller_id"`      // 商家ID
	Keyword       string `json:"keyword"`        // 模糊查询的关键字
	ServiceSn     string `json:"service_sn"`     // 售后服务单号
	OrderSn       string `json:"order_sn"`       // 订单编号
	GoodsName     string `json:"goods_name"`     // 商品名称
	RefundStatus  string `json:"refund_status"`  // 退款单状态 APPLY：申请，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
	RefundWay     string `json:"refund_way"`     // 退款方式 ORIGINAL：原路退回，OFFLINE：线下支付
	StartTime     int64  `json:"start_time"`     // 申请时间-起始时间
	EndTime       int64  `json:"end_time"`       // 申请时间-结束时间
	CreateChannel string `json:"create_channel"` // 退款单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
}
