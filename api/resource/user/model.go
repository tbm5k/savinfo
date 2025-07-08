package user

import (
	"github.com/google/uuid"
)

type UserDto struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

type User struct {
	ID uuid.UUID `gorm:"primarykey"`
    Name string
    Email string
}

type Users []*User

func (u *UserDto) ToModel() *User {
	return &User{
		Name: u.Name,
		Email: u.Email,
	}
}

func (u *User) ToDto() *UserDto {
	return &UserDto{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}

func (us Users) ToDtos() []*UserDto {
	usersDto := make([]*UserDto, len(us))
	for i, u := range us {
		usersDto[i] = u.ToDto()
	}
	return usersDto
}

