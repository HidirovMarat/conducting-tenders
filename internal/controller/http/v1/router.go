package v1

import (
	"conducting-tenders/internal/service"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}","uri":"${uri}", "status":${status},"error":"${error}"}` + "\n",
		Output: os.Stdout,
	}))
	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })

	v1 := handler.Group("/api")
	{
		newBidsRoutes(v1.Group("/bids"), services.Bid)
		newTendersRoutes(v1.Group("/tenders"), services.Tender)
		newPingRoutes(v1.Group("/ping"), services.Ping)
		newEmployeesRoutes(v1.Group("/employees"), services.Employee)
		newOrganizationsRoutes(v1.Group("/organizations"), services.Organization)
		newOrganizationResponsibleRoutes(v1.Group("/organizationResponsibles"), services.OrganizationResponsible)
	}
}
