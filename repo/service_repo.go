package repo

import (
	"context"
	"services/ent"
	"services/models"
)

type CatalogRepo interface {
	GetServices(ctx context.Context, req models.RequestParams) ([]*ent.Service, error)
	GetService(ctx context.Context, id int) (*ent.Service, error)
	GetCount(ctx context.Context, req models.RequestParams) (int, error)
}
