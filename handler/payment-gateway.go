package handler

import (
	"bitbucket.org/apps-for-bharat/gotools/blog"
	"epictectus/contract"
	"epictectus/domain"
	paymentGateway "epictectus/service/payment-gateway"
	"epictectus/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PgHandler struct {
	paymentGatewayService paymentGateway.PaymentGatewayService
}

func NewPaymentGatewayHandler(pgService paymentGateway.PaymentGatewayService) PgHandler {
	return PgHandler{paymentGatewayService: pgService}
}

func (h *PgHandler) CreateStandardPaymentLink(ctx *gin.Context) {
	requestContext := ctx.Request.Context()
	blog.InfoCtx(requestContext, "info-creating-standard-payment-link")
	var createStandardPaymentLinkRequest contract.CreateStandardPaymentLinkRequestRazorpay
	paymentProvider := ctx.GetHeader("x-payment-provider")
	if err := ctx.ShouldBindBodyWithJSON(&createStandardPaymentLinkRequest); err != nil {
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, createStandardPaymentLinkRequest.Validate(paymentProvider), "Invalid request body")
		ctx.JSON(httpStatus, errResp)
		return
	}
	requestContext = blog.SetValueInContext(requestContext, "paymentProvider", paymentProvider)
	err := h.paymentGatewayService.CreateStandardPaymentLink(requestContext, createStandardPaymentLinkRequest, domain.PaymentProvider(paymentProvider))
	if err != nil {
		ctx.JSON(utils.RenderError(err, "err-payment-link-creation-failed"))
		return
	}
	ctx.JSON(http.StatusOK, utils.RenderSuccess("info-payment-link-created-successfully"))
	return
}
