package v1

import (
	"conducting-tenders/internal/service"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type tendersRoutes struct {
	tendersService service.Tender
}

func newTendersRoutes(g *echo.Group, tendersService service.Tender) *tendersRoutes {
	r := &tendersRoutes{
		tendersService: tendersService,
	}

	g.POST("/new", r.createTender)
	g.GET("/my", r.getTendersByUsername)
	g.GET("", r.getTenders)
	g.GET("/{tenderId}/status", r.getTenderStatusById)
	g.PUT("/{tenderId}/edit", r.updateTenderStatusById)
	g.PATCH("/{tenderId}/edit", r.editTenderByIdAndUsername)
	//g.PUT("/tenders/{tenderId}/submit_decision", editSubmitDecisionById)
	//g.PUT("/tenders/{tenderId}/feedback", editFeelbackById)
	//g.PUT("/tenders/{tenderId}/rollback/{version}", editRollbackVersionById)
	//g.GET("/tenders/{tenderId}/reviews", getReviewsByTender)
	return r
}

func (r *tendersRoutes) createTender(c echo.Context) error {
	log.Println("tendersRoutes.createTender start")
	var input service.CreateTenderInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	tender, err := r.tendersService.CreateTender(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, tender)
}

func (r *tendersRoutes) getTendersByUsername(c echo.Context) error {
	var input service.GetTendersByUsernameInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	tenders, err := r.tendersService.GetTendersByUsername(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, tenders)
}

func (r *tendersRoutes) getTenders(c echo.Context) error {
	var input service.GetTendersInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	tenders, err := r.tendersService.GetTenders(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, tenders)
}

func (r *tendersRoutes) getTenderStatusById(c echo.Context) error {
	var input service.GetTenderStatusByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	status, err := r.tendersService.GetTenderStatusById(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, status)
}

func (r *tendersRoutes) updateTenderStatusById(c echo.Context) error {
	var input service.UpdateTenderStatusByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	tender, err := r.tendersService.UpdateTenderStatusById(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, tender)
}

func (r *tendersRoutes) editTenderByIdAndUsername(c echo.Context) error {
	var input service.EditTenderByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	tenders, err := r.tendersService.EditTenderById(c.Request().Context(), input)

	if err != nil {
		if errors.Is(err, service.ErrUserNotExistOrIncorrect) {
			return c.JSON(401, err.Error())
		}
		if errors.Is(err, service.ErrBidNotFind) {
			return c.JSON(404, service.ErrBidNotFind.Error())
		}
		if errors.Is(err, service.ErrInvalidRequestFormatOrParameters) {
			return c.JSON(400, err.Error())
		}
		if errors.Is(err, service.ErrNotEnoughRights) {
			return c.JSON(403, err.Error())
		}
		return c.JSON(404, err.Error())
	}

	return c.JSON(http.StatusOK, tenders)
}
