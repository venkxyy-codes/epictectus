package contract

type SignUpUser struct {
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (l *LoginUser) Validate() map[string]string {
	errors := make(map[string]string, 3)
	if len(l.Username) < 8 {
		errors["username"] = "err-username-should-not-be-lesser-than-8-characters"
	}
	return errors
}

func (c *SignUpUser) Validate() map[string]string {
	errors := make(map[string]string, 5)
	if c.Name == "" {
		errors["name"] = "err-name-is-required"
	}
	if len(c.Username) < 8 {
		errors["username"] = "err-username-should-not-be-lesser-than-8-characters"
	}
	return errors
}
