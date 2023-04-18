package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"services/ent"
	"services/models"
	"services/repo"
	"services/routes"
	"services/service"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Init() *echo.Echo {
	// Echo instance
	e := echo.New()
	client, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	catalogRepo := repo.ConfigureMySQLServiceRepo(client)
	serviceversionRepo := repo.ConfigureMySQLServiceVersionRepo(client)
	catalogSvc := service.ConfigureServiceImpl(catalogRepo, serviceversionRepo)
	v := validator.New()

	routes.InitServiceHandler(e, catalogSvc, v)
	return e
}

func ConnectDB() (*ent.Client, error) {
	dsn := "root:cashfree.123@tcp(localhost:3306)/test?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func TestMain(m *testing.M) {
	e := Init()
	// start server
	go func() {
		_ = e.Start("localhost:1323")
	}()
	fmt.Println("setup done")
	exitCode := m.Run()
	os.Exit(exitCode)
}

func callServicesAPI(t *testing.T, url string, resp interface{}) (int, interface{}) {

	res, err := http.Get(url)

	// validate no error and successful response
	assert.NoError(t, err)

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)

	err = json.Unmarshal(b, resp)
	if err != nil {
		assert.Fail(t, "Testing default url failed, ", url, err)
	}
	return res.StatusCode, resp
}

func TestDefaultServices(t *testing.T) {
	t.Log("Starting TestDefaultServices")
	// call endpoint using http
	url := "http://localhost:1323/services"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 4)
	assert.Equal(t, resp.Services[0].Id, 4)
	assert.Equal(t, resp.Services[1].Id, 3)
	assert.Equal(t, resp.Services[2].Id, 2)
	assert.Equal(t, resp.Services[3].Id, 1)
	t.Log("Ok")
}

func TestServicesLimit(t *testing.T) {
	t.Log("Starting TestServicesLimit")
	// call endpoint using http
	url := "http://localhost:1323/services?limit=2"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 2)
	assert.Equal(t, resp.Services[0].Id, 4)
	assert.Equal(t, resp.Services[1].Id, 3)
	t.Log("Ok")
}

func TestServicesStartAndLimit(t *testing.T) {
	t.Log("Starting TestServicesStartAndLimit")
	// call endpoint using http
	url := "http://localhost:1323/services?created_before=4&limit=2"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 2)
	assert.Equal(t, resp.Services[0].Id, 3)
	assert.Equal(t, resp.Services[1].Id, 2)
	t.Log("Ok")
}

func TestServicesEndAndLimit(t *testing.T) {
	t.Log("Starting TestServicesEndAndLimit")
	// call endpoint using http
	url := "http://localhost:1323/services?created_after=2&limit=2"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 2)
	assert.Equal(t, resp.Services[0].Id, 4)
	assert.Equal(t, resp.Services[1].Id, 3)
	t.Log("Ok")
}

func TestServicesFilterBy(t *testing.T) {
	t.Log("Starting TestServicesFilterBy")
	// call endpoint using http
	url := "http://localhost:1323/services?filter_by=chargeback"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 1)
	assert.Equal(t, resp.Services[0].Id, 3)
	t.Log("Ok")
}

func TestServicesFilterByDescription(t *testing.T) {
	t.Log("Starting TestServicesFilterBy")
	// call endpoint using http
	url := "http://localhost:1323/services?filter_by=test%20service"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 1)
	assert.Equal(t, resp.Services[0].Id, 1)
	t.Log("Ok")
}

func TestServicesFilterByDescriptionAndCreatedAfter(t *testing.T) {
	t.Log("Starting TestServicesFilterBy")
	// call endpoint using http
	url := "http://localhost:1323/services?filter_by=test%20service&created_after=1"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 0)
	t.Log("Ok")
}

func TestCreatedOnOrder(t *testing.T) {
	t.Log("Starting TestCreatedOnOrder")
	// call endpoint using http
	url := "http://localhost:1323/services"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 4)
	previousTime := time.Now()
	for _, s := range resp.Services {
		if s.CreatedOn.After(previousTime) {
			assert.Fail(t, "CreatedOn order not preserved")
		}
		previousTime = s.CreatedOn
	}
	t.Log("Ok")
}

func TestCreatedOnOrderWithCreatedBefore(t *testing.T) {
	t.Log("Starting TestCreatedOnOrderWithCreatedBefore")
	// call endpoint using http
	url := "http://localhost:1323/services?created_before=5"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, len(resp.Services), 4)
	previousTime := time.Now()
	for _, s := range resp.Services {
		if s.CreatedOn.After(previousTime) {
			assert.Fail(t, "CreatedOn order not preserved")
		}
		previousTime = s.CreatedOn
	}
	t.Log("Ok")
}

func TestVersionCount(t *testing.T) {
	t.Log("Starting TestVersionCount")
	url := "http://localhost:1323/services"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	actualCountMap := map[int]int{
		1: 4,
		2: 2,
		3: 1,
		4: 1,
	}
	assert.Equal(t, len(resp.Services), 4)
	for _, s := range resp.Services {
		val, ok := actualCountMap[s.Id]
		if !ok || val != s.VersionCount {
			assert.Fail(t, "Version count does not match")
		}
	}
	t.Log("Ok")
}

func TestSizeResponse(t *testing.T) {
	t.Log("Starting TestSizeResponse")
	// call endpoint using http
	url := "http://localhost:1323/services?created_after=1&limit=1"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, resp.Size, 3)
	t.Log("Ok")
}

func TestRequestInResponse(t *testing.T) {
	t.Log("Starting TestRequestInResponse")
	// call endpoint using http
	url := "http://localhost:1323/services?created_after=1&limit=1"
	resp := &models.CatalogResponse{}
	_, respStruct := callServicesAPI(t, url, resp)
	resp = respStruct.(*models.CatalogResponse)

	assert.Equal(t, resp.Request.CreatedAfter, 1)
	assert.Equal(t, resp.Request.Limit, 1)
	t.Log("Ok")
}
