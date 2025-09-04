package domain

type PaymentProvider string

const (
	Razorpay PaymentProvider = "razorpay"
)

func GetPaymentProviders() []string {
	return []string{string(Razorpay)}
}
