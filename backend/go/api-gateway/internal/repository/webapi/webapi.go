package webapi

import "webscrapper/api-gateway/internal/repository/webapi/internal"

type WebApi struct {
	Internal *internal.Internal
}

func New() *WebApi {
	return &WebApi{
		Internal: internal.New(),
	}
}
