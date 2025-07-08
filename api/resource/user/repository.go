package user

import (
    "github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&User{})
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) List() (Users, error) {
    users := make(Users, 0)
    if err := r.db.Find(&users).Error; err != nil {
        return nil, err
    }

    return users, nil
}

func (r *UserRepository) Create(user *User) (*User, error) {
    if err := r.db.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

func (r *UserRepository) Read(id uuid.UUID) (*User, error) {
    user := &User{}
    if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}
