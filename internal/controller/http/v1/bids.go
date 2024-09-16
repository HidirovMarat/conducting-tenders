package v1

import (
	"conducting-tenders/internal/entity/statusBid"
	"conducting-tenders/internal/service"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bidsRoutes struct {
	bidsService service.Bid
}

func newBidsRoutes(g *echo.Group, bidsService service.Bid) *bidsRoutes {
	r := &bidsRoutes{
		bidsService: bidsService,
	}

	g.POST("/new", r.createBid)
	g.GET("/my", r.getBidsByUsername)
	g.GET("/:tenderId/list", r.getBidsOfTenderById)
	g.GET("/:bidId/status", r.getBidStatusById)
	g.PUT("/:bidId/status", r.updateBidStatusById)
	g.PATCH("/:bidId/edit", r.editBidByIdAndUsername)
	g.PUT("/:bidId/submit_decision", r.updateSubmitDecisionById)
	g.PUT("/:bidId/rollback/:version", r.updateVersionBid)
	return r
}

func (r *bidsRoutes) createBid(c echo.Context) error {
	var input service.CreateBidInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	bid, err := r.bidsService.CreateBid(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bid)
}

func (r *bidsRoutes) getBidsByUsername(c echo.Context) error {
	var input service.GetBidsByUsernameInput
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	bids, err := r.bidsService.GetBidsByUsername(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bids)
}

func (r *bidsRoutes) getBidsOfTenderById(c echo.Context) error {
	var input service.GetBidsByTenderIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	bids, err := r.bidsService.GetBidsByTenderId(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bids)
}

func (r *bidsRoutes) getBidStatusById(c echo.Context) error {
	var input service.GetBidStatusByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	status, err := r.bidsService.GetBidStatusById(c.Request().Context(), input)

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

func (r *bidsRoutes) updateBidStatusById(c echo.Context) error {
	var input service.UpdateBidStatusByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}
	input.Username = c.QueryParam("username")
	input.Status = statusBid.Status(c.QueryParam("status"))

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	bid, err := r.bidsService.UpdateBidStatusById(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bid)
}

func (r *bidsRoutes) editBidByIdAndUsername(c echo.Context) error {
	var input service.EditBidByIdInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}
	input.Username = c.QueryParam("username")

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	bids, err := r.bidsService.EditBidById(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bids)
}

func (r *bidsRoutes) updateVersionBid(c echo.Context) error {
	var input service.UpdateVersionBidInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}
	input.Username = c.QueryParam("username")

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	bids, err := r.bidsService.UpdateVersionBid(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bids)
}

func (r *bidsRoutes) updateSubmitDecisionById(c echo.Context) error {
	var input service.UpdateSubmitDecisionInput
	err := c.Bind(&input)

	if err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}
	input.Username = c.QueryParam("username")
	input.Decision = c.QueryParam("decision")

	if err = c.Validate(input); err != nil {
		c.JSON(400, map[string]interface{}{"reason": err.Error()})
		return err
	}

	bids, err := r.bidsService.UpdateSubmitDecision(c.Request().Context(), input)

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

	return c.JSON(http.StatusOK, bids)

}
