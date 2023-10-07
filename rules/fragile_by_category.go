package rules

import "order-test/entity"

const homeAppliance = "eletrodoméstico"

func FragileByCategory(category string, o *entity.Order) {
	if category == homeAppliance {
		o.AddLabel(entity.FragileLabel)
	}
}
