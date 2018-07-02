package core

import (
	"code.int.thoseyears.com/golang/ppgo/engine"
	"code.int.thoseyears.com/golang/ppgo/helpers"
)

type Component struct {
	engine.BaseComponent
}

func (this *Component) Create(cfg interface{}) (interface{}, error) {
	c := &Config{}
	if err := helpers.StructDecode(cfg, c, "json"); err != nil {
		return nil, err
	}
	return New(c)
}

func (this *Component) Start(_, instance interface{}) error {
	m := instance.(*Manager)
	return m.Start()
}

func (this *Component) Stop(_, instance interface{}) error {
	m := instance.(*Manager)
	return m.Stop()
}
