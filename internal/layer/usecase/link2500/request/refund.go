package request

type Refund struct {
	RefundPrice float64 `json:"refund_price" validate:"required"`
	MerchantID  string  `json:"merchant_id" validate:"required"`
}
