package router

import (
	"github.com/c1emon/lemontree/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"sync"
)

var e *echo.Echo
var lock = &sync.Mutex{}

func SingletonEchoFactory() *echo.Echo {
	lock.Lock()
	defer lock.Unlock()
	if e == nil {
		e = echo.New()
		//if l, ok := e.Logger.(*log.Logger); ok {
		//	l.SetHeader("${time_rfc3339} ${level}")
		//	logBridge := &log.LogBridge{}
		//	l.SetOutput(logBridge)
		//}
		e.Logger = log.GetEchoLogrusLogger()
		//e.HTTPErrorHandler = nil

		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
				log.GetLogger().WithFields(logrus.Fields{
					"URI":    values.URI,
					"status": values.Status,
				}).Info("request")

				return nil
			},
		}))
	}

	return e
}

func BuildRouter() error {

	return nil
}
