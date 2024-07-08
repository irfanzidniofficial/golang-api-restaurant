package resto

import (
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/model/constant"
	"golang-api-restaurant/internal/respository/menu"
	"golang-api-restaurant/internal/respository/order"

	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (r *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

// Order(request model.OrderMenuRequest)(model.Order, error)
// GetOrderInfo(request model.GerOrderInfoRequest) (model.Order, error)

func (r *restoUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}
		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: int64(menuData.Price) * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}
	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProccessed,
		ProductOrders: productOrderData,
	}
	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, nil
	}
	return createOrderData, nil
}

func (r *restoUsecase) GetOrderInfo(request model.GerOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)
	if err != nil {
		return orderData, err
	}
	return orderData, nil
}
