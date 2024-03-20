package controller

import "github.com/oooiik/test_09.03.2024/internal/service"

type Interface interface {
	SetService(p service.Interface)
}

type controller struct {
	service service.Interface
}

func (c *controller) SetService(s service.Interface) {
	c.service = s
}
