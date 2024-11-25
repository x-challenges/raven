package sentry

import (
	"go.uber.org/zap"
)

type FxErrorHandler struct{}

func NewFxErrorHandler() *FxErrorHandler {
	return &FxErrorHandler{}
}

func (h *FxErrorHandler) HandleError(e error) {
	logger.Fatal("not handled error", zap.Error(e))
}
