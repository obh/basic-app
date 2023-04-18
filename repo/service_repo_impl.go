package repo

import (
	"context"
	"services/ent"
	"services/ent/service"
	"services/models"

	_ "github.com/go-sql-driver/mysql"
)

type catalogRepoMysql struct {
	entClient *ent.Client
}

func ConfigureMySQLServiceRepo(c *ent.Client) CatalogRepo {
	return &catalogRepoMysql{entClient: c}
}

// can be optimized further
func (c *catalogRepoMysql) GetServices(ctx context.Context, req models.RequestParams) ([]*ent.Service, error) {
	q := c.entClient.Debug().Service.Query()

	q = addWhereClause(q, &req)
	q = addFilterBy(q, &req)
	q = addOrderBy(q, &req)
	q = q.Limit(int(req.Limit))

	services, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (c *catalogRepoMysql) GetService(ctx context.Context, id int) (*ent.Service, error) {
	services, err := c.entClient.Service.Query().
		Where(service.IDEQ(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (c *catalogRepoMysql) GetCount(ctx context.Context, req models.RequestParams) (int, error) {
	q := c.entClient.Service.Query()
	q = addWhereClause(q, &req)
	q = addFilterBy(q, &req)
	count, err := q.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func addWhereClause(q *ent.ServiceQuery, req *models.RequestParams) *ent.ServiceQuery {
	if req.CreatedAfter > 0 {
		q = q.Where(service.IDGT(int(req.CreatedAfter)))
	} else if req.CreatedBefore > 0 {
		q = q.Where(service.IDLT(int(req.CreatedBefore)))
	}
	return q
}

func addOrderBy(q *ent.ServiceQuery, req *models.RequestParams) *ent.ServiceQuery {
	return q.Order(
		ent.Desc(service.FieldCreatedOn),
		ent.Desc(service.FieldID),
	)
}

func addFilterBy(q *ent.ServiceQuery, req *models.RequestParams) *ent.ServiceQuery {
	if req.FilterBy == "" {
		return q
	}
	return q.Where(
		service.Or(
			service.NameContains(req.FilterBy),
			service.DescriptionContains(req.FilterBy),
		),
	)
}
