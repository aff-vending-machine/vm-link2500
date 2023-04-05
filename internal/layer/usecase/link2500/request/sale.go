package request

type Sale struct {
	Price      float64 `json:"price" validate:"required"`
	MerchantID string  `json:"merchant_id" validate:"required"`
}
