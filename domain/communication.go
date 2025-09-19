package domain

type CommunicationChannels string

const (
	Whatsapp CommunicationChannels = "whatsapp"
)

type WhatsappProvider string

const (
	Angoor WhatsappProvider = "angoor"
)

func GetWhatsappProviders() []string {
	return []string{string(Angoor)}
}
