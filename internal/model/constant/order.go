package constant

import "golang-api-restaurant/internal/model"

const (
	OrderStatusProccessed model.OrderStatus = "proccessed"
	OrderStatusFinished   model.OrderStatus = "finished"
	OrderStatusFailed     model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinished  model.ProductOrderStatus = "finished"
)
