package repository

import (
	"context"
	"github.com/jenson/county/core/models"
	"github.com/jenson/county/county/global"
)

type CountyRepo struct {
}

func (cr CountyRepo) FindDetailByEditor(ctx context.Context, editor int32) (*models.CountyEntity, error) {
	var detail models.CountyEntity
	err := global.DB.Model(&models.CountyEntity{}).Where("editor = ?", editor).First(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}
