package routes

import (
	"net/http"
	"services/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func getErrorJSON(c echo.Context, err error) error {
	log.Info("getErrorJSON:: responding with err:", err.Error())
	switch err.(type) {
	case models.BadRequest:
		be := err.(models.BadRequest)
		if be.ErrType == models.ErrBadRequestParams {
			return c.JSON(http.StatusBadRequest,
				models.ResponseError{Title: models.TitleRequestInvalid, Detail: be.Detail},
			)
		} else if be.ErrType == models.ErrBadRequest {
			return c.JSON(http.StatusBadRequest,
				models.ResponseError{Title: models.TitleRequestInvalid, Detail: be.Detail},
			)
		}
	case models.ResourceNotFound:
		be := err.(models.ResourceNotFound)
		if be.ErrType == models.ErrBadRequestServiceId {
			return c.JSON(http.StatusNotFound,
				models.ResponseError{Title: models.TitleResourceNotFound, Detail: be.Detail},
			)
		}
	}
	return c.JSON(http.StatusInternalServerError,
		models.ResponseError{Title: models.TitleInternalServer, Detail: models.ErrUnknown})
}
