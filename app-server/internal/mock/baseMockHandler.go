package mock

import "app-server/internal/repository"

type BaseMockHandler struct {
	repository.BaseHandler
}

func NewBaseMockHandler(basehandler *repository.BaseHandler) *BaseMockHandler {
	return &BaseMockHandler{BaseHandler: *basehandler}
}
