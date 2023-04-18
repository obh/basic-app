package repo

import (
	"context"
	"services/ent"
	"services/models"
)

type ServiceVersions interface {
	GetVersionsForService(ctx context.Context, serviceId int) ([]*ent.ServiceVersion, error)
	GetVersionCountForServices(ctx context.Context, id []int) ([]models.ServiceVersionCount, error)
}
