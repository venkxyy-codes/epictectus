package contract

import (
	"epictectus/domain"
	"epictectus/utils"
	"slices"
	"strings"
)

type CreateStandardPaymentLinkRequestRazorpay struct {
	Amount          int64             `json:"amount"`
	Currency        string            `json:"currency,omitempty"`
	ReferenceID     string            `json:"reference_id,omitempty"`
	Description     string            `json:"description,omitempty"`
	ExpireBy        int64             `json:"expire_by,omitempty"`
	Notes           map[string]string `json:"notes,omitempty"`
	CustomerDetails struct {
		Name    string `json:"name"`
		Contact string `json:"contact"`
		Email   string `json:"email"`
	} `json:"customer,omitempty"`
	Notify struct {
		Sms   string `json:"sms,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"notify,omitempty"`
	AcceptPartial  bool   `json:"accept_partial,omitempty"`
	CallbackUrl    string `json:"callback_url,omitempty"`
	CallbackMethod string `json:"callback_method,omitempty"`
}

// Validate performs all request validations for a given payment provider.
func (r *CreateStandardPaymentLinkRequestRazorpay) Validate(paymentProvider string) []utils.ValidationError {
	errs := []utils.ValidationError{}

	// payment provider check (must be present and valid)
	if paymentProvider == "" || !slices.Contains(domain.GetPaymentProviders(), paymentProvider) {
		errs = append(errs, utils.ValidationError{
			Field:   "payment_provider",
			Message: "err-empty-or-invalid-payment-provider",
		})
	}

	// amount
	if r.Amount <= 0 {
		errs = append(errs, utils.ValidationError{
			Field:   "amount",
			Message: "err-amount-must-be-positive",
		})
	}

	// currency (optional on input; if provided, must be known)
	// Keep permissive: Razorpay supports multiple currencies. Enforce format & common code.
	// If you want to force INR only, replace with: if r.Currency != "" && r.Currency != "INR" { ... }
	if r.Currency != "" && len(r.Currency) != 3 {
		errs = append(errs, utils.ValidationError{
			Field:   "currency",
			Message: "err-currency-must-be-iso-4217-3-letter",
		})
	}

	// basic customer sanity (optional fields; format checks can be tightened later)
	if r.CustomerDetails.Email != "" && !strings.Contains(r.CustomerDetails.Email, "@") {
		errs = append(errs, utils.ValidationError{
			Field:   "customer_email",
			Message: "err-email-format-invalid",
		})
	}
	return errs
}
