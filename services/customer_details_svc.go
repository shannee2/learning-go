package services

import (
	"context"
	"learning-design/dao"
	"learning-design/models"
)

type CustomerDetailsService interface {
	GetRestaurantDetails(ctx context.Context, id int64) (*models.CustomerDetails, error)
	GetCustomerDetailsById(ctx context.Context, id int64) (*models.CustomerDetails, error)
}

func NewCustomerDetailsService(ctx context.Context, dao dao.Idao) CustomerDetailsService {
	return &CustomerDetailsServiceImpl{
		dao: dao,
	}
}

type CustomerDetailsServiceImpl struct {
	dao dao.Idao
}

func (cdi *CustomerDetailsServiceImpl) GetRestaurantDetails(ctx context.Context, id int64) (*models.CustomerDetails, error) {
	return nil, nil
}

func (cdi *CustomerDetailsServiceImpl) GetCustomerDetailsById(ctx context.Context, id int64) (*models.CustomerDetails, error) {
	name := cdi.dao.GetCustomerDetailsFromSource(ctx, id)
	return &models.CustomerDetails{Name: name}, nil
}
