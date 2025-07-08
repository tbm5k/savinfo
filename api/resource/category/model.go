package category

import (
	"github.com/google/uuid"
	"github.com/tbm5k/tss/api/resource/product"
)

type CategoryDto struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
	Products []product.ProductDto `json:"products"`
}

type Category struct {
	ID uuid.UUID `gorm:"primarykey"`
    Name string
	Products []product.Product
}

type Categories []*Category

func (c *CategoryDto) ToModel() *Category {
	return &Category{
		Name: c.Name,
	}
}

func (c *Category) ToDto() *CategoryDto {
	products := make([]product.ProductDto, len(c.Products))

	for i, p := range c.Products {
		products[i] = *p.ToDto() // dereference to assign value
	}

	return &CategoryDto{
		ID:       c.ID,
		Name:     c.Name,
		Products: products,
	}
}

func (cs Categories) ToDtos() []*CategoryDto {
	categoriesDto := make([]*CategoryDto, len(cs))
	for i, c := range cs {
		categoriesDto[i] = c.ToDto()
	}
	return categoriesDto
}

