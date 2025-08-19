package domain

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	PurchaseID int64
	Items      []ShippingItem
}
