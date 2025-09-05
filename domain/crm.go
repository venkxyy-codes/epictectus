package domain

import "strconv"

type CrmProvider string

type LeadsquaredActivityEventCode int64

const (
	Leadsquared CrmProvider = "leadsquared"

	RazorpayPaymentEvent    LeadsquaredActivityEventCode = 203
	PostRazorpayPaymentLink LeadsquaredActivityEventCode = 214
	GetPaymentRazorpayLink  LeadsquaredActivityEventCode = 216
)

func GetCrmProviders() []string {
	return []string{string(Leadsquared)}
}

func GetValidActivityEventCodesLeadsquared() []string {
	return []string{
		strconv.FormatInt(int64(PostRazorpayPaymentLink), 10),
		strconv.FormatInt(int64(GetPaymentRazorpayLink), 10),
		strconv.FormatInt(int64(RazorpayPaymentEvent), 10),
	}
}
