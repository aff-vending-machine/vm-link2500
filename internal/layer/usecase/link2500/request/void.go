package request

type Void struct {
	InvoiceNumber string `json:"invoice_number" validate:"required"`
}
