package usecase

import (
	"order-test/entity"
	"order-test/rules"
)

type CreateOrderDto struct {
	Product entity.Product
	Method  string
}

func OrderCreate(orderDto CreateOrderDto) entity.Order {

	payment := entity.Payment{
		Method: orderDto.Method,
		Value:  orderDto.Product.Value,
	}

	order := entity.Order{
		Product: orderDto.Product,
		Payment: payment,
	}

	rules.ValueToFreeShipping(orderDto.Product.Value, &order)

	return order
}
