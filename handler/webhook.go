package handler

import (
	"epictectus/blog"
	"epictectus/contract"
	webhookProcessor "epictectus/service/webhook_processor"
	"epictectus/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebhookHandler struct {
	webhookProcessorService webhookProcessor.WebhookProcessorService
}

func NewWebhookHandler(webhookProcessorService webhookProcessor.WebhookProcessorService) WebhookHandler {
	return WebhookHandler{webhookProcessorService: webhookProcessorService}
}

func (h *WebhookHandler) ProcessLeadsquaredActivityWebhook(ctx *gin.Context) {
	requestContext := ctx.Request.Context()
	blog.InfoCtx(requestContext, "info-processing-leadsquared-activity-webhook")
	var leadsquaredActivityWebhookRequest contract.LeadsquaredActivityWebhook
	if err := ctx.ShouldBindBodyWithJSON(&leadsquaredActivityWebhookRequest); err != nil {
		httpStatus, errResp := utils.RenderError(errors.ErrUnsupported, leadsquaredActivityWebhookRequest.Validate(), "Invalid request body")
		ctx.JSON(httpStatus, errResp)
		return
	}
	h.webhookProcessorService.HandleLeadsquaredWebhook(ctx, leadsquaredActivityWebhookRequest)
	ctx.JSON(http.StatusOK, utils.RenderSuccess("info-webhook-received-successfully"))
	return
}
