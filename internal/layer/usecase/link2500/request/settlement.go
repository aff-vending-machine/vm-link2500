package request

type Settlement struct {
	MerchantID string `json:"merchant_id" validate:"required"`
}
