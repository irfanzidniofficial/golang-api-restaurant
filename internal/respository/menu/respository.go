package menu

import (
	"context"
	"golang-api-restaurant/internal/model"
)

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
}
