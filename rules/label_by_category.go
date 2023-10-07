package rules

import "order-test/entity"

func LabelByCategory(categoryExpected string, label string, o *entity.Order) {
	if o.Product.Category == categoryExpected {
		o.AddLabel(label)
	}
}
