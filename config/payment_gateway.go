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

func GetRazorpayCredentials() RazorpayCredentials {
	return RazorpayCredentials{
		KeyId:     os.Getenv("RAZORPAY_KEY_ID"),
		KeySecret: os.Getenv("RAZORPAY_KEY_SECRET"),
	}
}

func GetRazorpayHttpConfig() RazorpayHttpConfig {
	return RazorpayHttpConfig{
		Host:                      os.Getenv("RAZORPAY_HOST"),
		CreatePaymentLinkEndpoint: os.Getenv("RAZORPAY_CREATE_PAYMENT_LINK_ENDPOINT"),
	}
}
