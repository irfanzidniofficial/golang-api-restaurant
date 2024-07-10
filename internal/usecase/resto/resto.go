package resto

import (
	"errors"
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/model/constant"
	"golang-api-restaurant/internal/respository/menu"
	"golang-api-restaurant/internal/respository/order"
	"golang-api-restaurant/internal/respository/user"

	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

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
		ReferenceID:   request.ReferenceID,
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

func (r *restoUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	userRegisted, err := r.userRepo.CheckRegistered(request.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRegisted {
		return model.User{}, errors.New("user already registered")
	}
	userHash, err := r.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, nil
	}
	userData, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.NewString(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}
