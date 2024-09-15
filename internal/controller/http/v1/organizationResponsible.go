package v1

import (
	"conducting-tenders/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type organizationResponsibleRoutes struct {
	oganizationResponsibleService service.OrganizationResponsible
}

func newOrganizationResponsibleRoutes(g *echo.Group, oganizationResponsibleService service.OrganizationResponsible) *organizationResponsibleRoutes {
	r := &organizationResponsibleRoutes{
		oganizationResponsibleService: oganizationResponsibleService,
	}

	g.GET("", r.getOrganizationResponsibles)
	return r
}

func (r *organizationResponsibleRoutes) getOrganizationResponsibles(c echo.Context) error {
	var input service.GetOrganizationResponsiblesInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	oganizationResponsible, err := r.oganizationResponsibleService.GetOrganizationResponsibles(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, oganizationResponsible)
}
