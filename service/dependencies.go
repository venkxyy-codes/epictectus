package service

import (
	"epictectus/appcontext"
	_ "epictectus/appcontext"
	"epictectus/clients"
	"epictectus/repo"
	"epictectus/service/communication"
	"epictectus/service/crm"
	"epictectus/service/payment-gateway"
	"epictectus/service/user"
	webhookProcessor "epictectus/service/webhook_processor"
)

type ServerDependencies struct {
	UserService             user.UserService
	PaymentGatewayService   payment_gateway.PaymentGatewayService
	CrmService              crm.CrmService
	WebhookProcessorService webhookProcessor.WebhookProcessorService
}

func InstantiateServerDependencies() *ServerDependencies {
	dbClient := appcontext.GetDBClient()
	userRepo := repo.NewUserRepository(dbClient)
	userServ := user.NewUserService(userRepo)
	baseClient := clients.NewBaseClient()
	crmService := crm.NewCrmService(baseClient)
	commService := communication.NewCommService(baseClient)
	paymentGatewayService := payment_gateway.NewPaymentGatewayService(crmService, commService, baseClient)
	webhookProcessorService := webhookProcessor.NewWebhookProcessorService(crmService, paymentGatewayService)
	return &ServerDependencies{
		UserService:             userServ,
		PaymentGatewayService:   paymentGatewayService,
		CrmService:              crmService,
		WebhookProcessorService: webhookProcessorService,
	}
}
