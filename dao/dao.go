package dao

import (
	"context"
	"sync"
)

type Idao interface {
	GetCustomerDetailsFromSource(ctx context.Context, id int64) string
}

var daoSync sync.Once
var dao Idao

func NewDAO(ctx context.Context) Idao {
	daoSync.Do(func() {
		dao = &DAOImpl{}
	})
	return dao
}

type DAOImpl struct {
}

func (di *DAOImpl) GetCustomerDetailsFromSource(ctx context.Context, id int64) string {
	return "Shannee"
}
