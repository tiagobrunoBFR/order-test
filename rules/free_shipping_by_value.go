package rules

import "order-test/entity"

const freeShippingValue = 1000

func FreeShippingByValue(value int, o *entity.Order) {
	if value > freeShippingValue {
		o.AddLabel(entity.LabelFreeShipping)
	}
}
