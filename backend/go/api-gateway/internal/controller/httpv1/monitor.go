package httpv1

import "github.com/gin-gonic/gin"

type Monitor interface {
	Ping(ctx *gin.Context)
	InitRoutes(r *gin.RouterGroup)
}

type monitor struct{}

func NewMonitor() Monitor {
	return &monitor{}
}

func (h *monitor) InitRoutes(r *gin.RouterGroup) {
	m := r.Group("/monitor/")
	{
		m.GET("/ping", h.Ping)
	}
}

func (h *monitor) Ping(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
