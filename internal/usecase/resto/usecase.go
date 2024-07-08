package resto

import "golang-api-restaurant/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
