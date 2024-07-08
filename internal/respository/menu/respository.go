package menu

import "golang-api-restaurant/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
