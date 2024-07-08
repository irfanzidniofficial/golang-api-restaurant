package menu

import "golang-api-restaurant/internal/model"

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string)(model.MenuItem, error)
}
