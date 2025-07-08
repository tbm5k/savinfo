package order

import (
	"time"

	"github.com/google/uuid"
)

type OrderDto struct {
    ID uuid.UUID `json:"id"`
    Total uint `json:"total"`
    Status string `json:"status"`
	OrderDate time.Time `json:"orderDate"`
	UserID uuid.UUID `json:"userId"`
}

type Order struct {
	ID uuid.UUID `gorm:"primarykey"`
    Total uint
    Status string
	OrderDate time.Time
	UserID uuid.UUID `gorm:"foreignKey:UserID;references:ID"`
}

type Orders []*Order

func (o *OrderDto) ToModel() *Order {
	return &Order{
		Total: o.Total,
		Status: o.Status,
		OrderDate: o.OrderDate,
		UserID: o.UserID,
	}
}

func (o *Order) ToDto() *OrderDto {
	return &OrderDto{
		ID: o.ID,
		Total: o.Total,
		Status: o.Status,
		OrderDate: o.OrderDate,
		UserID: o.UserID,
	}
}

func (ols Orders) ToDtos() []*OrderDto {
	ordersDto := make([]*OrderDto, len(ols))
	for i, p := range ols {
		ordersDto[i] = p.ToDto()
	}
	return ordersDto
}

