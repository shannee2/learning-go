package api

import (
	"learning-design/models"
	"learning-design/services"
)

type ICustomerAPI interface {
	GetCustomerNameById(id int64) (*models.CustomerDetails, error)
}

func NewCustomerAPI(custSvc services.CustomerDetailsService) ICustomerAPI {
	return &customerAPIImpl{
		custSvc: custSvc,
	}
}

type customerAPIImpl struct {
	custSvc services.CustomerDetailsService
}

func (custAPiImpl *customerAPIImpl) GetCustomerNameById(id int64) (*models.CustomerDetails, error) {
	return custAPiImpl.custSvc.GetCustomerNameById(id)
}
