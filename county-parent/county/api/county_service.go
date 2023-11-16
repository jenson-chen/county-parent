package api

import (
	"context"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/county/repository"
	"github.com/jenson/county/county/service"
	"log"
)

type CountyService struct {
	service.UnimplementedCountyRpcServer
}

func (cs *CountyService) Insert(ctx context.Context, in *models.CountyEntity) (*models.NormalResult, error) {
	return nil, nil
}

func (cs *CountyService) QueryDetailByEditor(ctx context.Context, in *models.EditorRequest) (*models.CountyEntity, error) {
	repo := repository.RepositoryInstance.CountyRepository
	county, err := repo.FindDetailByEditor(ctx, in.GetEditor())
	if err != nil {
		log.Printf("database query error,%#v", err)
		return nil, err
	}
	return county, nil
}

func (cs *CountyService) DeleteByEditor(ctx context.Context, in *models.EditorRequest) (*models.NormalResult, error) {
	return nil, nil
}

func (cs *CountyService) UpdateByEditor(ctx context.Context, in *models.UpdateRequest) (*models.NormalResult, error) {
	return nil, nil
}
