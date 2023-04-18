package repo

import (
	"context"
	"fmt"
	"services/ent"
	"services/ent/serviceversion"
	"services/models"

	"entgo.io/ent/dialect/sql"
)

type versionRepo struct {
	entClient *ent.Client
}

func ConfigureMySQLServiceVersionRepo(c *ent.Client) ServiceVersions {
	return &versionRepo{entClient: c}
}

func (c *versionRepo) GetVersionsForService(ctx context.Context, serviceId int) ([]*ent.ServiceVersion, error) {
	versions, err := c.entClient.ServiceVersion.Query().
		Where(serviceversion.ServiceIDEQ(serviceId)).
		All(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	fmt.Println("found versions for service:", serviceId, versions)
	return versions, nil
}

func (c *versionRepo) GetVersionCountForServices(ctx context.Context, ids []int) ([]models.ServiceVersionCount, error) {
	var v []models.ServiceVersionCount
	err := c.entClient.ServiceVersion.Query().
		Where(
			func(s *sql.Selector) {
				s.Where(sql.InInts(serviceversion.FieldServiceID, ids...))
			}).
		GroupBy(serviceversion.ServiceColumn).
		Aggregate(ent.Count()).
		Scan(ctx, &v)
	if err != nil {
		fmt.Println("failed querying services: ", err)
		return nil, err
	}
	fmt.Println("Found versions --> ", v)
	return v, nil
}

/*
//get all data using sql
func (c *versionRepo) testMethod(ctx context.Context, ids []int) {
	type versionCount struct {
		Id    int `json:"id"`
		Count int `json:"count"`
	}
	var v []versionCount
	c.entClient.Service.Query().
		Order(ent.Desc(service.FieldCreatedOn)).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.InInts(service.FieldID, ids...))
			}).
		GroupBy(service.FieldID).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(serviceversion.Table)
			s.Join(t).On(s.C(service.FieldID), t.C(serviceversion.FieldServiceID))
			return sql.As(sql.Count(t.C(serviceversion.FieldID)), "count")
		}).
		Scan(ctx, &v)
	fmt.Println(v)
}
*/
