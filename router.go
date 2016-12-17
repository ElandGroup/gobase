package main

import (
	"goBasis/service"
	"net/http"

	"github.com/labstack/echo"
)

func RegApi() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello gobasis")
	})

	api := e.Group("/api")
	v1 := api.Group("/v1")
	basis := v1.Group("/basis")
	basis.GET("", service.GetBasis)
	flowcontrol := v1.Group("/flowcontrol")
	flowcontrol.GET("/if", service.Getif)
	flowcontrol.GET("/for", service.Getfor)
	flowcontrol.GET("/switch", service.Getswitch)
	funcApi := v1.Group("/func")
	funcApi.GET("", service.Getfunc)
	funcApi.GET("/pass", service.GetFuncPass)

}
