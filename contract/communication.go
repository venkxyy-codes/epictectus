package contract

type SendPaymentLinkToCustomer struct {
	LeadsquaredProspectId string  `json:"prospect_id"`
	PaymentLink           string  `json:"payment_link"`
	PaymentAmount         float64 `json:"payment_amount"`
	CustomerPhoneNumber   string  `json:"customer_phone_number"`
}
