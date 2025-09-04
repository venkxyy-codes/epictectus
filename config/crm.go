package config

import "os"

type LeadsquaredHttpConfig struct {
	Host                              string `env:"LSQ_HOST"`
	FetchLeadUsingPhoneNumberEndpoint string `env:"LSQ_FETCH_LEAD_USING_PHONE_NUMBER_ENDPOINT"`
	PostActivityToLead                string `env:"LSQ_POST_ACTIVITY_TO_LEAD"`
}

type LeadsquaredCredentials struct {
	AccessKey    string `env:"LSQ_ACCESS_KEY"`
	SecretAccess string `env:"LSQ_SECRET_ACCESS"`
}

func GetLeadsquaredCredentials() LeadsquaredCredentials {
	return LeadsquaredCredentials{
		AccessKey:    os.Getenv("LSQ_ACCESS_KEY"),
		SecretAccess: os.Getenv("LSQ_SECRET_ACCESS"),
	}
}

func GetLeadsquaredHttpConfig() LeadsquaredHttpConfig {
	return LeadsquaredHttpConfig{
		Host:                              os.Getenv("LSQ_HOST"),
		FetchLeadUsingPhoneNumberEndpoint: os.Getenv("LSQ_FETCH_LEAD_USING_PHONE_NUMBER_ENDPOINT"),
		PostActivityToLead:                os.Getenv("LSQ_POST_ACTIVITY_TO_LEAD"),
	}
}
