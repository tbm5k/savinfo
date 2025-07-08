package product

import (
	"github.com/google/uuid"
)

type ProductDto struct {
    ID uuid.UUID `json:"id"`
    Title string `json:"title"`
    Price uint `json:"price"`
	CategoryID uuid.UUID `json:"categoryId"`
}

type Product struct {
	ID uuid.UUID `gorm:"primarykey"`
    Title string
    Price uint
	CategoryID uuid.UUID `gorm:"foreignKey:CategoryID;references:ID"`
}

type Products []*Product

func (p *ProductDto) ToModel() *Product {
	return &Product{
		Title: p.Title,
		Price: p.Price,
		CategoryID: p.CategoryID,
	}
}

func (p *Product) ToDto() *ProductDto {
	return &ProductDto{
		ID: p.ID,
		Title: p.Title,
		Price: p.Price,
		CategoryID: p.CategoryID,
	}
}

func (ps Products) toDtos() []*ProductDto {
	productsDto := make([]*ProductDto, len(ps))
	for i, p := range ps {
		productsDto[i] = p.ToDto()
	}
	return productsDto
}

