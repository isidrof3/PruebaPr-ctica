package endpoint

import (
	"context"
	translate "translator/endpoint/translate"
	"translator/process"
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	Translate endpoint.Endpoint
}

// @Summary Return Response
// @Description translate Service
// @Tags translate
// @Accept  json
// @Produce  json
// @Param Request body translate.Request true "Translate Request"
// @Success 200 {object} translate.Response
// @Router /translator/translate [post]
func MakeTranslateEndpoint(p process.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(translate.Request)
		resp, err := p.GetTranslateProcess(req)
		return resp, err
	}
}
