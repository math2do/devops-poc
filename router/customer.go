package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veltris/devops-poc/config"
	"github.com/veltris/devops-poc/dtos"
	"github.com/veltris/devops-poc/services/customer"
	"github.com/veltris/devops-poc/services/kubernetes"
)

// CustomerHandler ...
type CustomerHandler struct {
	log             *log.Logger
	customerService *customer.Service
}

// NewCustomerHandler ...
func NewCustomerHandler(log *log.Logger) *CustomerHandler {
	return &CustomerHandler{
		log:             log,
		customerService: customer.NewService(log, kubernetes.NewK8sClient(log)),
	}
}

// provisionVMs ...
func (ch *CustomerHandler) provisionVMs(ctx *gin.Context) {

	var req dtos.ProvisionVMRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		errorResponse(ctx, fmt.Errorf("unable to read body"), http.StatusBadGateway)
		return
	}

	// TODO: validate req payload

	res, err := ch.customerService.ProvisionVMs(&req)
	if err != nil {
		errorResponse(ctx, fmt.Errorf(config.InternalServerErr), http.StatusInternalServerError)
		return
	}

	successJSON(ctx, res)
}
