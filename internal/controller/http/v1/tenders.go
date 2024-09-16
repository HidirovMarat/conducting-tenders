package v1

import (
	"conducting-tenders/internal/service"
	"errors"
	"log"
	"net/http"
	"strconv"

	serviceType "conducting-tenders/internal/entity/service-type"
	"conducting-tenders/internal/entity/statusTender"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//validate:"dive,required,oneof=Construction Delivery Manufacture"` //validate:"dive,required,oneof=Construction Delivery Manufacture"

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
	tenderId := g.Group("/:tenderId")
	{
		tenderId.GET("/status", r.getTenderStatusById)
		tenderId.PUT("/status", r.updateTenderStatusById)
		tenderId.PATCH("/edit", r.editTenderByIdAndUsername)
	}

	//g.PUT("/tenders/{tenderId}/submit_decision", editSubmitDecisionById)
	//g.PUT("/tenders/{tenderId}/feedback", editFeelbackById)
	g.PUT("/:tenderId/rollback/:version", r.updateVersionTender)
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

	// Извлечение limit и offset
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	// Преобразуем limit и offset из строки в число
	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			input.Limit = l
		} else {
			return c.JSON(400, map[string]interface{}{"reason": "Invalid limit"})
		}
	}
	if offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			input.Offset = o
		} else {
			return c.JSON(400, map[string]interface{}{"reason": "Invalid offset"})
		}
	}

	// Извлечение serviceType
	serviceTypes := c.QueryParams()["service_type"]
	if len(serviceTypes) == 0 {
		return c.JSON(400, map[string]interface{}{"reason": "service_type is required"})
	}

	// Преобразуем serviceTypes в нужный тип []serviceType.ServiceType
	for _, s := range serviceTypes {
		// Здесь можно добавить логику для преобразования строки в serviceType.ServiceType
		switch s {
		case "Construction", "Delivery", "Manufacture":
			input.ServiceType = append(input.ServiceType, serviceType.ServiceType(s))
		default:
			return c.JSON(400, map[string]interface{}{"reason": "Invalid serviceType"})
		}
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
	log.Println(input)

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
		log.Println("Yesssssssssssssssss")
		return c.JSON(404, map[string]interface{}{"reason": err.Error()})
	}

	return c.JSON(http.StatusOK, status)
}

func (r *tendersRoutes) updateTenderStatusById(c echo.Context) error {
	var input service.UpdateTenderStatusByIdInput
	//err := c.Bind(&input)
	username := c.QueryParam("username")
	status := c.QueryParam("status")
	tenderId := c.Param("tenderId")
	var err error

	input.Status = statusTender.Status(status)
	input.TenderId, err = uuid.Parse(tenderId)

	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}

	input.Username = username

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
	//input := new(service.EditTenderByIdInput)
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
	}
	input.Username = c.QueryParam("username")

	if err = c.Validate(input); err != nil {
		return c.JSON(400, map[string]interface{}{"reason": err.Error()})
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

func (r *tendersRoutes) updateVersionTender(c echo.Context) error {
	var input service.UpdateVersionTenderInput
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

	tenders, err := r.tendersService.UpdateVersionTender(c.Request().Context(), input)

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
