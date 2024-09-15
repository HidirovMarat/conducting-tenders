package v1

import (
	"conducting-tenders/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type employeesRoutes struct {
	employeesService service.Employee
}

func newEmployeesRoutes(g *echo.Group, employeesService service.Employee) *employeesRoutes {
	r := &employeesRoutes{
		employeesService: employeesService,
	}

	g.GET("", r.getEmployees)
	return r
}

func (r *employeesRoutes) getEmployees(c echo.Context) error {
	var input service.GetEmployeesInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	employees, err := r.employeesService.GetEmployees(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrEmployeeNotFind) {
			return c.JSON(404, service.ErrEmployeeNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, employees)
}
