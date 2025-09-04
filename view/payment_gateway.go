package view

type RazorpayCreatePaymentLinkResponse struct {
	AcceptPartial  bool   `json:"accept_partial"`
	Amount         int    `json:"amount"`
	AmountPaid     int    `json:"amount_paid"`
	CallbackMethod string `json:"callback_method"`
	CallbackUrl    string `json:"callback_url"`
	CancelledAt    int    `json:"cancelled_at"`
	CreatedAt      int    `json:"created_at"`
	Currency       string `json:"currency"`
	Customer       struct {
		Contact string `json:"contact"`
		Email   string `json:"email"`
		Name    string `json:"name"`
	} `json:"customer"`
	Description           string      `json:"description"`
	ExpireBy              int         `json:"expire_by"`
	ExpiredAt             int         `json:"expired_at"`
	FirstMinPartialAmount int         `json:"first_min_partial_amount"`
	Id                    string      `json:"id"`
	Notes                 interface{} `json:"notes"`
	Notify                struct {
		Email    bool `json:"email"`
		Sms      bool `json:"sms"`
		Whatsapp bool `json:"whatsapp"`
	} `json:"notify"`
	Payments       interface{}   `json:"payments"`
	ReferenceId    string        `json:"reference_id"`
	ReminderEnable bool          `json:"reminder_enable"`
	Reminders      []interface{} `json:"reminders"`
	ShortUrl       string        `json:"short_url"`
	Status         string        `json:"status"`
	UpdatedAt      int           `json:"updated_at"`
	UpiLink        bool          `json:"upi_link"`
	UserId         string        `json:"user_id"`
	WhatsappLink   bool          `json:"whatsapp_link"`
}
