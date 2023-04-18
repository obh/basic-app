package repo

import (
	"context"
	"services/ent"
	"services/models"
)

type ServicesRepo interface {
	GetAll(ctx context.Context, req models.RequestParams) ([]*ent.Service, error)
	GetById(ctx context.Context, id int) (*ent.Service, error)
	GetCount(ctx context.Context, req models.RequestParams) (int, error)
}
