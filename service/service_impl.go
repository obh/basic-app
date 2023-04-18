package service

import (
	"context"
	"errors"
	"services/ent"
	"services/models"
	"services/repo"
)

type catalogServiceImpl struct {
	catalogRepo repo.CatalogRepo
	versionRepo repo.ServiceVersions
}

func ConfigureServiceImpl(r repo.CatalogRepo, s repo.ServiceVersions) CatalogService {
	return &catalogServiceImpl{catalogRepo: r, versionRepo: s}
}

func (c *catalogServiceImpl) GetServices(ctx context.Context, req models.RequestParams) (*models.CatalogResponse, error) {
	services, err := c.catalogRepo.GetServices(ctx, req)
	if err != nil {
		return nil, errors.New(models.ErrInternalServerError)
	}
	ids := []int{}
	for _, s := range services {
		ids = append(ids, s.ID)
	}
	counts, err := c.versionRepo.GetVersionCountForServices(ctx, ids)
	if err != nil {
		return nil, errors.New(models.ErrInternalServerError)
	}
	itemMap := make(map[int]int)
	for _, item := range counts {
		itemMap[item.ServiceId] = item.Count
	}
	svcJson := serviceMapper(services, itemMap)
	size, err := c.catalogRepo.GetCount(ctx, req)
	if err != nil {
		return nil, errors.New(models.ErrInternalServerError)
	}
	resp := &models.CatalogResponse{Services: svcJson, Size: size, Count: len(services)}
	return resp, nil
}

func (c *catalogServiceImpl) GetService(ctx context.Context, id int) (*models.Service, error) {
	s, err := c.catalogRepo.GetService(ctx, id)
	if err != nil {
		return nil, errors.New(models.ErrInternalServerError)
	}
	v, err := c.versionRepo.GetVersionsForService(ctx, id)
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
