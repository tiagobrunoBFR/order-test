package entity

const LabelFreeShipping = "frete-gratis"
const LabelFragile = "frágil"
const LabelGift = "presente"

const PaymentMethodPix = "Pix"
const PaymentMethodBankSlip = "Boleto"

const CategoryElectronic = "eletrônico"
const CategoryHomeAppliance = "eletrodoméstico"
const CategoryChildren = "infantil"

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
