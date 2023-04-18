package service

import (
	"context"
	"services/ent"
	"services/models"
	"services/repo"
)

type serviceControllerImpl struct {
	servicesRepo        repo.ServicesRepo
	serviceVersionsRepo repo.ServiceVersions
}

func ConfigureServiceController(r repo.ServicesRepo, s repo.ServiceVersions) ServiceController {
	return &serviceControllerImpl{servicesRepo: r, serviceVersionsRepo: s}
}

func (c *serviceControllerImpl) GetServices(ctx context.Context, req models.RequestParams) (*models.Services, error) {
	services, err := c.servicesRepo.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	ids := []int{}
	for _, s := range services {
		ids = append(ids, s.ID)
	}
	counts, err := c.serviceVersionsRepo.GetServiceVersionCounts(ctx, ids)
	if err != nil {
		return nil, err
	}
	itemMap := make(map[int]int)
	for _, item := range counts {
		itemMap[item.ServiceId] = item.Count
	}
	svcJson := serviceMapper(services, itemMap)
	size, err := c.servicesRepo.GetCount(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &models.Services{Items: svcJson, Size: size, Count: len(services)}
	return resp, nil
}

func (c *serviceControllerImpl) GetService(ctx context.Context, id int) (*models.Service, error) {
	s, err := c.servicesRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	v, err := c.serviceVersionsRepo.GetAll(ctx, id)
	return responseMapper(s, v), nil
}

func serviceMapper(services []*ent.Service, versions map[int]int) []models.Service {
	svcJson := []models.Service{}
	for _, s := range services {
		if v, ok := versions[s.ID]; ok {
			t := models.Service{
				Id:          s.ID,
				Name:        s.Name,
				Description: s.Description,
				CreatedOn:   s.CreatedOn,
			}
			t.VersionCount = v
			svcJson = append(svcJson, t)
		}
	}
	return svcJson
}

func responseMapper(service *ent.Service, versions []*ent.ServiceVersion) *models.Service {
	versionJSON := []models.ServiceVersion{}
	for _, v := range versions {
		versionJSON = append(versionJSON, models.ServiceVersion{
			Id:        v.ID,
			Name:      v.Name,
			CreatedOn: v.CreatedOn,
		})
	}
	return &models.Service{
		Id:           service.ID,
		Name:         service.Name,
		Description:  service.Description,
		CreatedOn:    service.CreatedOn,
		VersionCount: len(versionJSON),
		Versions:     versionJSON,
	}
}
