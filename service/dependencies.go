package service

import (
	"epictectus/appcontext"
	_ "epictectus/appcontext"
	"epictectus/repo"
	"epictectus/service/crm"
	"epictectus/service/payment-gateway"
	"epictectus/service/user"
)

type ServerDependencies struct {
	UserService           user.UserService
	PaymentGatewayService payment_gateway.PaymentGatewayService
}

func InstantiateServerDependencies() *ServerDependencies {
	dbClient := appcontext.GetDBClient()
	userRepo := repo.NewUserRepository(dbClient)
	userServ := user.NewUserService(userRepo)
	lsqService := crm.NewLeadsquaredService()
	paymentGatewayService := payment_gateway.NewPaymentGatewayService(lsqService)
	return &ServerDependencies{
		UserService:           userServ,
		PaymentGatewayService: paymentGatewayService,
	}
}
