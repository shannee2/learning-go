package main

import (
	"context"
	"fmt"
	apis "learning-design/api"
	dao "learning-design/dao"
	"learning-design/services"
)

func main() {
	ctx := context.Background()
	daobj := dao.NewDAO(ctx)
	svc := services.NewCustomerDetailsService(ctx, daobj)
	api := apis.NewCustomerAPI(ctx, svc)
	name, _ := api.GetCustomerDetailsById(ctx, 1)
	fmt.Println(name.Name)
}
