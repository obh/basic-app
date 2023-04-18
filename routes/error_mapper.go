package routes

import (
	"fmt"
	"net/http"
	"services/models"

	"github.com/labstack/echo/v4"
)

func getErrorJSON(c echo.Context, err error) error {
	fmt.Println("here with : ", err)
	switch err.(type) {
	case models.BadRequest:
		be := err.(models.BadRequest)
		if be.ErrType == models.ErrBadRequestParams {
			return c.JSON(http.StatusBadRequest,
				models.ResponseError{Title: models.TitleRequestInvalid, Detail: be.Detail},
			)
		} else if be.ErrType == models.ErrBadRequestServiceId {
			return c.JSON(http.StatusBadRequest,
				models.ResponseError{Title: models.TitleResourceNotFound, Detail: be.Detail},
			)
		} else if be.ErrType == models.ErrBadRequest {
			return c.JSON(http.StatusBadRequest,
				models.ResponseError{Title: models.TitleRequestInvalid, Detail: be.Detail},
			)
		}
	}
	return c.JSON(http.StatusInternalServerError,
		models.ResponseError{Title: models.TitleInternalServer, Detail: models.ErrUnknown})
}
