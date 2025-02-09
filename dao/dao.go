package dao

type Idao interface {
	GetCustomerDetailsFromSource(id int64) string
}

func NewDAO() Idao {
	return &DAOImpl{}
}

type DAOImpl struct {
}

func (di *DAOImpl) GetCustomerDetailsFromSource(id int64) string {
	return "Shannee"
}
