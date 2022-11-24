package main

import (
	order_api "github.com/Cesarmosqueira/coffeshop_api/cmd/backend/order/api"
	product_api "github.com/Cesarmosqueira/coffeshop_api/cmd/backend/product/api"
	handler "github.com/Cesarmosqueira/coffeshop_api/pkg/handler"
)

func main() {
	server := handler.NewServer()
	server.EnableApi(product_api.NewProductWebApi)
	server.EnableApi(order_api.NewOrderWebApi)

	server.Run()
}
