package config

import "os"

type RazorpayHttpConfig struct {
	Host                      string `env:"RAZORPAY_HOST"`
	CreatePaymentLinkEndpoint string `env:"RAZORPAY_CREATE_PAYMENT_LINK_ENDPOINT"`
}

type RazorpayCredentials struct {
	KeyId     string `env:"RAZORPAY_KEY_ID"`
	KeySecret string `env:"RAZORPAY_KEY_SECRET"`
}

func GetRazorpayCredentials() map[string]interface{} {
	return map[string]interface{}{
		"key_id":     os.Getenv("RAZORPAY_KEY_ID"),
		"key_secret": os.Getenv("RAZORPAY_KEY_SECRET"),
	}
}

func GetRazorpayHttpConfig() map[string]interface{} {
	return map[string]interface{}{
		"host":                         os.Getenv("RAZORPAY_HOST"),
		"create_payment_link_endpoint": os.Getenv("RAZORPAY_CREATE_PAYMENT_LINK_ENDPOINT"),
	}
}
