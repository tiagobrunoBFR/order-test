package rules

import "order-test/entity"

func FragileByCategory(category string, o *entity.Order) {
	if category == entity.CategoryHomeAppliance {
		o.AddLabel(entity.LabelFragile)
	}

	if category == entity.CategoryChildren {
		o.AddLabel(entity.LabelGift)
	}
}
