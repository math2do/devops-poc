package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veltris/devops-poc/services/kubernetes"
)

// NewServer ...
func NewServer() *gin.Engine {
	// init configs, or singletons
	err := kubernetes.InitKubernetesClient()
	if err != nil {
		log.Fatalf("unable to init k8s client")
		return nil
	}

	router := gin.Default()

	v1 := router.Group("/v1/automation")

	v1.GET("/health", func(ctx *gin.Context) {
		successJSON(ctx, map[string]string{"status": "healthy"})
		return
	})

	// customer related APIs
	{
		customer := v1.Group("/customer")
		addAutomationRoutes(customer)
	}

	return router
}

// addAutomationRoutes ...
func addAutomationRoutes(router *gin.RouterGroup) {
	customerHandler := NewCustomerHandler(log.Default())

	router.POST("/create-vm", customerHandler.provisionVMs)
}

func successJSON(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.JSON(http.StatusOK, data)
}

func errorResponse(ctx *gin.Context, err error, statusCode int) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.JSON(statusCode, gin.H{
		"status": "FAILED",
		"error":  err.Error(),
	})
}
