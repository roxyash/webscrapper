package httpv1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webscrapper/api-gateway/internal/model"
	"webscrapper/api-gateway/internal/service"
)

type Auth interface {
	InitRoutes(handler *gin.RouterGroup)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type auth struct {
	s service.Auth
}

func NewAuth(s service.Auth) Auth {
	return &auth{s: s}
}

func (h *auth) InitRoutes(handler *gin.RouterGroup) {
	a := handler.Group("/auth/")
	{
		a.POST("/login", h.Login)
		a.POST("/register", h.Register)
	}
}

func (h *auth) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid body request",
		})
		return
	}

	resp, err := h.s.Login(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *auth) Register(ctx *gin.Context) {
	var req model.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid body request",
		})
		return
	}

	resp, err := h.s.Register(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
