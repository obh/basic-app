package test

import (
	"net/http"
	"regexp"
	"services/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidLimit(t *testing.T) {
	t.Log("Starting TestInvalidLimit")
	url := "http://localhost:1323/services?limit=-1000"
	resp := &models.ResponseError{}
	statusCode, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.ResponseError)

	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, resp.Title, models.TitleRequestInvalid)
	t.Log("Ok")
}

func TestInvalidFilterByCharacters(t *testing.T) {
	t.Log("Starting TestInvalidFilterByCharacters")
	url := "http://localhost:1323/services?filter_by=<script>"
	resp := &models.ResponseError{}
	statusCode, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.ResponseError)

	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, resp.Title, models.TitleRequestInvalid)
	t.Log("Ok")
}

func TestRegex(t *testing.T) {
	t.Log("Starting TestRegex")
	match, err := regexp.MatchString(`^[a-zA-Z0-9&_ ,.\-()\/]+$`, "<script>")
	assert.Nil(t, err)
	assert.Equal(t, match, false)
	t.Log("Ok")
}

func TestBothCreatedAfterAndBefore(t *testing.T) {
	t.Log("Starting TestBothCreatedAfterAndBefore")
	url := "http://localhost:1323/services?created_after=2&created_before=4"
	resp := &models.ResponseError{}
	statusCode, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.ResponseError)

	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, resp.Title, models.TitleRequestInvalid)
	t.Log("Ok")
}
