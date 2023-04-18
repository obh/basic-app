package service

import (
	"context"
	"services/models"
)

type CatalogService interface {
	GetServices(c context.Context, req models.RequestParams) (*models.CatalogResponse, error)
	GetService(c context.Context, id int) (*models.Service, error)
}
