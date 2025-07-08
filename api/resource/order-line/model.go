package orderline

import (
	"github.com/google/uuid"
	"github.com/tbm5k/tss/api/resource/order"
	"github.com/tbm5k/tss/api/resource/product"
)

type Form struct {
    Quantity uint `json:"quantity"`
	ProductID string `json:"productId"`
}

type OrderLineDto struct {
    ID uuid.UUID `json:"id"`
    UnitPrice uint `json:"unitPrice"`
    Quantity uint `json:"quantity"`
	OrderID uuid.UUID `json:"orderId"`
	ProductID uuid.UUID `json:"productId"`
	Order order.OrderDto `json:"order"`
	Product product.ProductDto `json:"product"`
}

type OrderLine struct {
	ID uuid.UUID `gorm:"primarykey"`
    UnitPrice uint
    Quantity uint
	OrderID uuid.UUID `gorm:"foreignKey:OrderID;references:ID"`
	ProductID uuid.UUID `gorm:"foreignKey:ProductID;references:ID"`

	Order order.Order
	Product product.Product
}

type OrderLines []*OrderLine

func (ol *OrderLineDto) ToModel() *OrderLine {
	return &OrderLine{
		UnitPrice: ol.UnitPrice,
		Quantity: ol.Quantity,
		OrderID: ol.OrderID,
		ProductID: ol.ProductID,
	}
}

func (ol *OrderLine) ToDto() *OrderLineDto {
	return &OrderLineDto{
		ID: ol.ID,
		UnitPrice: ol.UnitPrice,
		Quantity: ol.Quantity,
		Order: *ol.Order.ToDto(),
		Product: *ol.Product.ToDto(),
	}
}

func (ps OrderLines) ToDtos() []*OrderLineDto {
	orderLinesDto := make([]*OrderLineDto, len(ps))
	for i, p := range ps {
		orderLinesDto[i] = p.ToDto()
	}
	return orderLinesDto
}

