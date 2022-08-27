package domain

import "gorm.io/gorm"

// All type definition related to the domain is defined here
type User struct {
	gorm.Model
	Name string `json:"name"`
}

// UserUseCase here is the interface that will be implemented by the use case
type UserUseCase interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(u *User) (*User, error)
	Update(id int, u *User) (*User, error)
	Delete(id int) error
}

// UserRepository here is the interface that will be implemented by the repository
type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(u *User) (*User, error)
	Update(id int, u *User) (*User, error)
	Delete(id int) error
}
