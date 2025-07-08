package orderline

import (
    "github.com/google/uuid"

	"gorm.io/gorm"
)

type OrderLineRepository struct {
    db *gorm.DB
}

func NewOrderLineRepository(db *gorm.DB) *OrderLineRepository {
	db.AutoMigrate(&OrderLine{})
    return &OrderLineRepository{
        db: db,
    }
}

func (r *OrderLineRepository) List() (OrderLines, error) {
    orderLines := make(OrderLines, 0)
    if err := r.db.Preload("Order").Preload("Product").Find(&orderLines).Error; err != nil {
        return nil, err
    }

    return orderLines, nil
}

func (r *OrderLineRepository) Create(orderLine *OrderLine) (*OrderLine, error) {
    if err := r.db.Create(orderLine).Error; err != nil {
        return nil, err
    }

    return orderLine, nil
}

func (r *OrderLineRepository) Read(id uuid.UUID) (*OrderLine, error) {
    orderLine := &OrderLine{}
    if err := r.db.Preload("Order", "Product").Where("id = ?", id).First(orderLine).Error; err != nil {
        return nil, err
    }

    return orderLine, nil
}
