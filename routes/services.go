package routes

import (
	"errors"
	"fmt"
	"html"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"services/models"
	"services/service"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const createdAfterQuery = "created_after"
const createdBeforeQuery = "created_before"
const filterByQuery = "filter_by"
const limitQuery = "limit"
const serviceIdPathParam = "id"
const defaultFrom = 1
const DEFAULT_LIMIT = 10
const MAX_LIMIT = 50

// /services?from=1&count=10&filterBy=test&orderBy=created_on&orderType=desc

type serviceHandler struct {
	serviceController service.ServiceController
	validate          *validator.Validate
}

func InitServiceHandler(e *echo.Echo, svc service.ServiceController, v *validator.Validate) error {
	s := &serviceHandler{serviceController: svc, validate: v}

	e.Add("GET", "/services", s.getServices)
	e.Add("GET", "/services/:id", s.getService)

	return nil
}

func (h *serviceHandler) getServices(c echo.Context) error {
	req, err := h.getReqParams(c)
	if err != nil {
		fmt.Println(err)
		return getErrorJSON(c,
			models.BadRequest{BaseError: models.BaseError{ErrType: models.ErrBadRequest, Detail: err.Error()}})
	}
	log.Info("verified request, going forward")
	resp, err := h.serviceController.GetServices(c.Request().Context(), *req)
	resp.Request = *req
	if err != nil {
		fmt.Println(err)
		return getErrorJSON(c, models.InternalServerError{BaseError: models.BaseError{
			ErrType: models.ErrInternalServerError, Detail: err.Error()}})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *serviceHandler) getService(c echo.Context) error {
	id := c.Param(serviceIdPathParam)
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return getErrorJSON(c,
			models.BadRequest{BaseError: models.BaseError{ErrType: models.ErrBadRequest, Detail: err.Error()}})
	}
	s, err := h.serviceController.GetService(c.Request().Context(), int(idInt))
	if s == nil {
		return getErrorJSON(c, models.ResourceNotFound{BaseError: models.BaseError{
			ErrType: models.ErrBadRequestServiceId, Detail: models.ErrBadRequestServiceId}})
	}
	return c.JSON(http.StatusOK, s)
}

func (h *serviceHandler) getReqParams(c echo.Context) (*models.RequestParams, error) {
	createdAfter, err := extractIntParam(createdAfterQuery, 0, c.QueryParams())
	if err != nil {
		return nil, err
	}
	createdBefore, err := extractIntParam(createdBeforeQuery, 0, c.QueryParams())
	if err != nil {
		return nil, err
	}
	limit, err := extractIntParam(limitQuery, DEFAULT_LIMIT, c.QueryParams())
	if err != nil {
		return nil, err
	}
	limit = int(math.Min(float64(MAX_LIMIT), float64(limit)))

	r := &models.RequestParams{
		CreatedAfter:  createdAfter,
		CreatedBefore: createdBefore,
		Limit:         limit,
		FilterBy:      html.EscapeString(c.QueryParams().Get(filterByQuery)),
	}
	err = validate(r, h.validate)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func extractIntParam(param string, defaultVal int, queryParams url.Values) (int, error) {
	if v := queryParams.Get(param); v != "" {
		if i, err := strconv.Atoi(v); err != nil {
			return 0, err
		} else {
			return i, nil
		}
	}
	return defaultVal, nil
}

func validate(r *models.RequestParams, validate *validator.Validate) error {
	fmt.Println("calling validate: ", r)
	match, err := regexp.MatchString(`^[a-zA-Z0-9&_ ,.\-()\/]+$`, r.FilterBy)
	if err != nil {
		return err
	}
	if !match && r.FilterBy != "" {
		return errors.New("filter_by has invalid characters")
	}
	if r.CreatedAfter > 0 && r.CreatedBefore > 0 {
		return errors.New("Cannot set both created_after and created_before in same query")
	}
	return validate.Struct(r)
}
