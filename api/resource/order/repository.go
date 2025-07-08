package order

import (
    "github.com/google/uuid"

	"gorm.io/gorm"
)

type OrderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	db.AutoMigrate(&Order{})
    return &OrderRepository{
        db: db,
    }
}

func (r *OrderRepository) GetByStatus(status string) (*Order, error) {
    var order Order
    if err := r.db.Where("status = ?", status).First(&order).Error; err != nil {
        return nil, err
    }

    return &order, nil
}

func (r *OrderRepository) GetPendingOrder() (*Order, error) {
    var order Order
    if err := r.db.Where("status = ?", "pending").First(&order).Error; err != nil {
        return nil, err
    }

    return &order, nil
}

func (r *OrderRepository) List() (Orders, error) {
    orders := make(Orders, 0)
    if err := r.db.Find(&orders).Error; err != nil {
        return nil, err
    }

    return orders, nil
}

func (r *OrderRepository) Create(order *Order) (*Order, error) {
    if err := r.db.Create(order).Error; err != nil {
        return nil, err
    }

    return order, nil
}

func (r *OrderRepository) Update(id uuid.UUID, order *Order) error {
	res := r.db.Where("id = ?", id).Updates(order)
	if res.Error != nil {
		return res.Error
	}

    return nil
}

func (r *OrderRepository) Read(id uuid.UUID) (*Order, error) {
    order := &Order{}
    if err := r.db.Where("id = ?", id).First(order).Error; err != nil {
        return nil, err
    }

    return order, nil
}
