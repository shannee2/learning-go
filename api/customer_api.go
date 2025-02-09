package api

import (
	"context"
	"learning-design/models"
	"learning-design/services"
)

type ICustomerAPI interface {
	GetCustomerDetailsById(ctx context.Context, id int64) (*models.CustomerDetails, error)
}

func NewCustomerAPI(ctx context.Context, custSvc services.CustomerDetailsService) ICustomerAPI {
	return &customerAPIImpl{
		custSvc: custSvc,
	}
}

type customerAPIImpl struct {
	custSvc services.CustomerDetailsService
}

func (custAPiImpl *customerAPIImpl) GetCustomerDetailsById(ctx context.Context, id int64) (*models.CustomerDetails, error) {
	return custAPiImpl.custSvc.GetCustomerDetailsById(ctx, id)
}
