package main

import (
	"fmt"
	api2 "learning-design/api"
	dao2 "learning-design/dao"
	"learning-design/services"
)

func main() {
	dao := dao2.NewDAO()
	svc := services.NewCustomerDetailsService(dao)
	api := api2.NewCustomerAPI(svc)
	name, _ := api.GetCustomerNameById(1)
	fmt.Println(name.Name)
}
