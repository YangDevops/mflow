package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mflow/apps/pipeline"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service pipeline.Service
	log     logger.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(pipeline.AppName)
	h.service = ioc.GetController(pipeline.AppName).(pipeline.Service)
	return nil
}

func (h *handler) Name() string {
	return pipeline.AppName
}

func (h *handler) Version() string {
	return "v1"
}
