package entity

const FreeShippingLabel = "frete-gratis"

type Product struct {
	Category string
	Value    int
}

type Payment struct {
	Method string
	Value  int
}

type Order struct {
	Product Product
	Payment Payment
	Labels  []string
}

func (o *Order) AddLabel(label string) {
	o.Labels = append(o.Labels, label)
}
