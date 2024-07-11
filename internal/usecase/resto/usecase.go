package resto

import "golang-api-restaurant/internal/model"

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	Order(request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request model.GerOrderInfoRequest) (model.Order, error)
	RegisterUser(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.UserSession, error)
}
