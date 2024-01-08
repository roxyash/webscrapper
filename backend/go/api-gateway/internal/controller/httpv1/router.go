package httpv1

import (
	"github.com/gin-gonic/gin"
	"webscrapper/api-gateway/internal/service"
)

const (
	BasePath = "/api/v1"
)

type Handler interface {
	InitRoutes(handler *gin.RouterGroup)
}

type Router struct {
	*gin.Engine
}

// New method for create router
func New(service *service.Service) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// Create router
	r := &Router{engine}

	// Init routes
	r.InitRoutes(service)

	return r
}

func (r *Router) InitRoutes(service *service.Service) {
	handlers := []Handler{
		NewAuth(service.Auth),
		NewMonitor(),
	}

	// API Group
	v1 := gin.New()

	// Init blueprints
	for _, h := range handlers {
		h.InitRoutes(v1.Group(BasePath))
	}
}
