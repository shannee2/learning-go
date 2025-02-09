package services

import (
	"learning-design/dao"
	"learning-design/models"
)

type CustomerDetailsService interface {
	GetRestaurantDetails(id int64) (*models.CustomerDetails, error)
	GetCustomerNameById(id int64) (*models.CustomerDetails, error)
}

func NewCustomerDetailsService(dao dao.Idao) CustomerDetailsService {
	return &CustomerDetailsServiceImpl{
		dao: dao,
	}
}

type CustomerDetailsServiceImpl struct {
	dao dao.Idao
}

func (cdi *CustomerDetailsServiceImpl) GetRestaurantDetails(id int64) (*models.CustomerDetails, error) {
	return nil, nil
}

func (cdi *CustomerDetailsServiceImpl) GetCustomerNameById(id int64) (*models.CustomerDetails, error) {
	name := cdi.dao.GetCustomerDetailsFromSource(id)
	return &models.CustomerDetails{Name: name}, nil
}
