package v1

import (
	"conducting-tenders/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type organizationsRoutes struct {
	organizationsService service.Organization
}

func newOrganizationsRoutes(g *echo.Group, organizationsService service.Organization) *organizationsRoutes {
	r := &organizationsRoutes{
		organizationsService: organizationsService,
	}

	g.GET("", r.getOrganizations)
	return r
}

func (r *organizationsRoutes) getOrganizations(c echo.Context) error {
	var input service.GetOrganizationsInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	organizations, err := r.organizationsService.GetOrganizations(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrOrganizationNotFind) {
			return c.JSON(404, service.ErrOrganizationNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, organizations)
}
