package service

import (
	"context"
	"services/models"
)

type ServiceController interface {
	GetServices(c context.Context, req models.RequestParams) (*models.Services, error)
	GetService(c context.Context, id int) (*models.Service, error)
}
