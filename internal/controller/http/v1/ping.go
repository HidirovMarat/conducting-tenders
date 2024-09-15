package v1

import (
	"conducting-tenders/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type pingRoutes struct {
	pingService service.Ping
}

func newPingRoutes(g *echo.Group, pingService service.Ping) *pingRoutes {
	r := &pingRoutes{
		pingService: pingService,
	}

	g.GET("", r.ping)

	return r
}

func (r *pingRoutes) ping(c echo.Context) error {

	err := r.pingService.Ping(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "ok")
}
