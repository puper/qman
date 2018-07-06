package bootstrap

import (
	"code.int.thoseyears.com/golang/ppgo/components"
	"code.int.thoseyears.com/golang/ppgo/engine"
	"github.com/puper/qman/producer/app"
	"github.com/puper/qman/producer/components/producer"
	"github.com/puper/qman/producer/routes"
	"github.com/spf13/viper"
)

func Bootstrap(configFile string) error {
	conf := viper.New()
	conf.SetConfigFile(configFile)
	if err := conf.ReadInConfig(); err != nil {
		return err
	}
	myapp := app.Create(&engine.Config{*conf})
	myapp.RegisterComponents(components.Components())
	myapp.RegisterComponent("producer", &producer.Component{})
	if err := myapp.Init(); err != nil {
		return err
	}
	routes.Configure(app.GetServer())
	if err := myapp.Start(); err != nil {
		return err
	}
	return nil
}
