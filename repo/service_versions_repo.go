package repo

import (
	"context"
	"services/ent"
	"services/models"
)

type ServiceVersions interface {
	GetAll(ctx context.Context, serviceId int) ([]*ent.ServiceVersion, error)
	GetServiceVersionCounts(ctx context.Context, id []int) ([]models.ServiceVersionsCount, error)
}
