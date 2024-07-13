package resto

import (
	"context"
	"errors"
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/model/constant"
	"golang-api-restaurant/internal/respository/menu"
	"golang-api-restaurant/internal/respository/order"
	"golang-api-restaurant/internal/respository/user"
	"golang-api-restaurant/internal/tracing"

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

func (r *restoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()
	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (r *restoUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(ctx, orderProduct.OrderCode)
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
		UserID:        request.UserID,
		Status:        constant.OrderStatusProccessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}
	createOrderData, err := r.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		return model.Order{}, nil
	}
	return createOrderData, nil
}

func (r *restoUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()
	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return orderData, err
	}
	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("unauthorized")
	}
	return orderData, nil
}

func (r *restoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	userRegisted, err := r.userRepo.CheckRegistered(ctx, request.Username)

	if err != nil {
		return model.User{}, err
	}

	if userRegisted {
		return model.User{}, errors.New("user already registered")
	}
	userHash, err := r.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		return model.User{}, nil
	}
	userData, err := r.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.NewString(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *restoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()
	userData, err := r.userRepo.GetUserData(ctx, request.Username)
	if err != nil {
		return model.UserSession{}, err
	}
	verified, err := r.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}
	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}
	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}
	return userSession, nil
}

func (r *restoUsecase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()
	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", nil
	}
	return userID, nil
}
