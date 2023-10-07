package rules

import "order-test/entity"

const freeShippingValue = 1000

func ValueToFreeShipping(value int, o *entity.Order) {
	if value > freeShippingValue {
		o.AddLabel(entity.FreeShippingLabel)
	}
}
