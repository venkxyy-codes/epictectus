package handler

import (
	"epictectus/blog"
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
	var createStandardPaymentLinkRequest contract.CreateStandardPaymentLink
	paymentProvider := ctx.GetHeader("x-payment-provider")
	if err := ctx.ShouldBindBodyWithJSON(&createStandardPaymentLinkRequest); err != nil {
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, createStandardPaymentLinkRequest.Validate(paymentProvider), "Invalid request body")
		ctx.JSON(httpStatus, errResp)
		return
	}
	requestContext = blog.SetValueInContext(requestContext, "paymentProvider", paymentProvider)
	switch paymentProvider {
	case string(domain.Razorpay):
		err := h.paymentGatewayService.CreateStandardPaymentLinkRazorpay(requestContext, createStandardPaymentLinkRequest, true, domain.Leadsquared, true, domain.Angoor)
		if err != nil {
			ctx.JSON(utils.RenderError(err, "err-payment-link-creation-failed"))
			return
		}
	default:
		err := h.paymentGatewayService.CreateStandardPaymentLinkRazorpay(requestContext, createStandardPaymentLinkRequest, true, domain.Leadsquared, true, domain.Angoor)
		if err != nil {
			ctx.JSON(utils.RenderError(err, "err-payment-link-creation-failed"))
			return
		}
	}
	ctx.JSON(http.StatusOK, utils.RenderSuccess("info-payment-link-created-successfully"))
	return
}
