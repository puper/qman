package app

import (
	"code.int.thoseyears.com/golang/ppgo/components/dbman"
	"code.int.thoseyears.com/golang/ppgo/components/log"
	"code.int.thoseyears.com/golang/ppgo/components/redis"
	"code.int.thoseyears.com/golang/ppgo/engine"
	"github.com/Sirupsen/logrus"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/puper/qman/consumer/core"
)

var (
	app *engine.Engine
)

func Create(cfg *engine.Config) *engine.Engine {
	app = engine.New(cfg)
	return app
}

func Get() *engine.Engine {
	return app
}

func GetDB() *dbman.DBMan {
	return app.GetInstance("db").(*dbman.DBMan)
}

func GetLog(name string) *logrus.Logger {
	return app.GetInstance("log").(*log.Log).Get(name)
}

func GetRedis(name string) redigo.Conn {
	return app.GetInstance("redis").(*redis.Redis).Get(name)
}

func GetCore() *core.Manager {
	return app.GetInstance("core").(*core.Manager)
}
