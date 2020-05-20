package adminserver

import (
	"net/http"

	"github.com/dsrhub/dsrhub/pkg/config"
	"github.com/dsrhub/dsrhub/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AdminServer struct {
	HostPort string

	echo *echo.Echo
	tx   *gorm.DB
	dsl  *model.DSL
}

func (as *AdminServer) Start() error {
	if err := as.prepare(); err != nil {
		return err
	}
	as.echo.Start(as.HostPort)
	return nil
}

func (as *AdminServer) Teardown() {
	return
}

func (as *AdminServer) prepare() error {
	as.tx = model.GetDB()
	as.prepareEcho()
	return as.prepareDSL()
}

func (as *AdminServer) prepareDSL() error {
	as.dsl = &model.DSL{}
	return as.dsl.Load(config.ENV.DSLPath)
}

func (as *AdminServer) prepareEcho() {
	as.echo = echo.New()
	as.echo.HideBanner = true
	as.echo.Use(middleware.Logger())

	g := as.echo.Group("/admin")
	g.GET("/workflows", as.funcGetWorkflows())
	g.GET("/dsl", as.funcGetDSL())
}

func (as *AdminServer) funcGetWorkflows() func(c echo.Context) error {
	return func(c echo.Context) error {
		return nil
	}
}

func (as *AdminServer) funcGetDSL() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, as.dsl.YAML())
	}
}
