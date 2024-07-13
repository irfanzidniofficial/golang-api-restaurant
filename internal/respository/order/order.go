package order

import (
	"context"
	"golang-api-restaurant/internal/model"
	"golang-api-restaurant/internal/tracking"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}

func (or *orderRepo) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	ctx, span := tracking.CreateSpan(ctx, "CreateOrder")
	defer span.End()
	if err := or.db.WithContext(ctx).Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (or *orderRepo) GetOrderInfo(ctx context.Context, orderID string) (model.Order, error) {
	ctx, span := tracking.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()
	var data model.Order

	if err := or.db.WithContext(ctx).Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
