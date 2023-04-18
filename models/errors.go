package models

type ResponseError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type BaseError struct {
	ErrType string
	Detail  string
}

func (b BaseError) Error() string {
	return b.ErrType
}

func (b BaseError) ErrorDetails() string {
	return b.Detail
}

type BadRequest struct {
	BaseError
}

type ResourceNotFound struct {
	BaseError
}

type InternalServerError struct {
	BaseError
}

var (
	TitleRequestInvalid   = "REQUEST_INVALID"
	TitleResourceNotFound = "RESOURCE_NOT_FOUND"
	TitleInternalServer   = "INTERNAL_SERVER_ERROR"

	ErrBadRequest          = "Bad Request"
	ErrInternalServerError = "Internal Server Error"
	ErrBadRequestServiceId = "Bad request, Service ID doesn't exist"
	ErrBadRequestParams    = "Bad request, Input validation failed"
	ErrUnknown             = "Unknown error occured - mostly some uncaptured scenario"
)
