package repo

import (
	"context"
	"services/ent"
	"services/ent/service"
	"services/ent/serviceversion"
	"services/models"

	"entgo.io/ent/dialect/sql"
	"github.com/labstack/gommon/log"
)

type mySQLServiceVersionsRepo struct {
	entClient *ent.Client
}

func ConfigureMySQLServiceVersionsRepo(c *ent.Client) ServiceVersions {
	return &mySQLServiceVersionsRepo{entClient: c}
}

func (c *mySQLServiceVersionsRepo) GetAll(ctx context.Context, serviceId int) ([]*ent.ServiceVersion, error) {
	versions, err := c.entClient.ServiceVersion.Query().
		Where(serviceversion.ServiceIDEQ(serviceId)).
		Order(ent.Desc(service.FieldCreatedOn)).
		All(ctx)
	if err != nil {
		log.Info("GetVersionsForService::Error querying versions of a ServiceId error: ", err, serviceId)
		return nil, err
	}
	log.Info("GetVersionsForService::found versions for service:", serviceId, versions)
	return versions, nil
}

func (c *mySQLServiceVersionsRepo) GetServiceVersionCounts(ctx context.Context, ids []int) ([]models.ServiceVersionsCount, error) {
	var v []models.ServiceVersionsCount
	err := c.entClient.ServiceVersion.Query().
		Where(
			func(s *sql.Selector) {
				s.Where(sql.InInts(serviceversion.FieldServiceID, ids...))
			}).
		GroupBy(serviceversion.ServiceColumn).
		Aggregate(ent.Count()).
		Scan(ctx, &v)
	if err != nil {
		log.Info("GetVersionCountForServices::failed querying services: ", err)
		return nil, err
	}
	return v, nil
}
