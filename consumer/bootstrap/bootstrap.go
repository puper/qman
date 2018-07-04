package bootstrap

import (
	"code.int.thoseyears.com/golang/ppgo/components"
	"code.int.thoseyears.com/golang/ppgo/engine"
	"github.com/spf13/viper"

	"github.com/puper/qman/consumer/app"
	"github.com/puper/qman/consumer/core"
	"github.com/puper/qman/consumer/storage/mysql"
)

func Bootstrap(configFile string) error {
	conf := viper.New()
	conf.SetConfigFile(configFile)
	if err := conf.ReadInConfig(); err != nil {
		return err
	}
	myapp := app.Create(&engine.Config{*conf})
	myapp.RegisterComponents(components.Components())
	myapp.RegisterComponent("core", &core.Component{})
	if err := myapp.Init(); err != nil {
		return err
	}
	app.GetCore().SetStorage(&mysql.Storage{})
	if err := myapp.Start(); err != nil {
		return err
	}
	return nil
}
