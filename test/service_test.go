package test

import (
	"net/http"
	"services/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultServiceRequest(t *testing.T) {
	t.Log("Starting TestDefaultServiceRequest")
	url := "http://localhost:1323/services/1"
	resp := &models.Service{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.Service)

	assert.Equal(t, resp.Id, 1)
	assert.Equal(t, resp.Name, "test")
	t.Log("Ok")
}

func TestVersionsInService(t *testing.T) {
	t.Log("Starting TestVersionsInService")
	url := "http://localhost:1323/services/2"
	resp := &models.Service{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.Service)

	actualVersions := []string{"service2.v2", "service2.v1"}
	assert.Equal(t, resp.VersionCount, 2)
	for i, s := range resp.Versions {
		if actualVersions[i] != s.Name {
			assert.Fail(t, "Versions do not match")
		}
	}
	t.Log("Ok")
}

func TestVersionsOrder(t *testing.T) {
	t.Log("Starting TestVersionsOrder")
	url := "http://localhost:1323/services/1"
	resp := &models.Service{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.Service)

	assert.Equal(t, resp.VersionCount, 4)
	previousTime := time.Now()
	for _, s := range resp.Versions {
		if s.CreatedOn.After(previousTime) {
			assert.Fail(t, "CreatedOn order not preserved")
		}
		previousTime = s.CreatedOn
	}
	t.Log("Ok")
}

func TestServiceIdNotFound(t *testing.T) {
	t.Log("Starting TestInvalidId")
	url := "http://localhost:1323/services/101"
	resp := &models.ResponseError{}
	statusCode, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.ResponseError)

	assert.Equal(t, statusCode, http.StatusNotFound)
	t.Log("Ok")
}

func TestServiceIdInvalid(t *testing.T) {
	t.Log("Starting TestInvalidId")
	url := "http://localhost:1323/services/<script>"
	resp := &models.ResponseError{}
	statusCode, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.ResponseError)

	assert.Equal(t, statusCode, http.StatusBadRequest)
	t.Log("Ok")
}
