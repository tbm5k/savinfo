package category

import (
    "github.com/google/uuid"

	"gorm.io/gorm"
)

type CategoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	db.AutoMigrate(&Category{})
    return &CategoryRepository{
        db: db,
    }
}

func (r *CategoryRepository) List() (Categories, error) {
    categorys := make(Categories, 0)
    if err := r.db.Preload("Products").Find(&categorys).Error; err != nil {
        return nil, err
    }

    return categorys, nil
}

func (r *CategoryRepository) Create(category *Category) (*Category, error) {
    if err := r.db.Create(category).Error; err != nil {
        return nil, err
    }

    return category, nil
}

func (r *CategoryRepository) Read(id uuid.UUID) (*Category, error) {
    category := &Category{}
    if err := r.db.Preload("Products").Where("id = ?", id).First(category).Error; err != nil {
        return nil, err
    }

    return category, nil
}

