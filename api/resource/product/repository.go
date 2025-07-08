package product

import (
	"log"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	db.AutoMigrate(&Product{})
    return &ProductRepository{
        db: db,
    }
}

func (p *ProductRepository) List() (Products, error) {
    products := make(Products, 0)
    if err := p.db.Find(&products).Error; err != nil {
        return nil, err
    }

    return products, nil
}

func (p *ProductRepository) GetByID(id uuid.UUID) (*Product, error) {
	product := &Product{}
	log.Printf("id %v \n", id)
    if err := p.db.Where("id = ?", id).First(product).Error; err != nil {
		log.Printf("Product lookup failed for id %v: %v", id, err)
        return nil, err
    }
    return product, nil
}

func (p *ProductRepository) Create(product *Product) (*Product, error) {
    if err := p.db.Create(product).Error; err != nil {
        return nil, err
    }

    return product, nil
}

func (p *ProductRepository) Read(id uuid.UUID) (*Product, error) {
    product := &Product{}
    if err := p.db.Where("id = ?", id).First(product).Error; err != nil {
        return nil, err
    }

    return product, nil
}
