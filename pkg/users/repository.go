package users

import (
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})

	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAll() ([]User, error) {
	var users []User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetById(id int) (User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *userRepository) Create(usr *User) (*User, error) {
	if err := u.db.Create(usr).Error; err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userRepository) Update(id int, newUsr *User) (*User, error) {
	var user User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	u.db.Model(&user).Updates(newUsr)

	return &user, nil

}

func (u *userRepository) Delete(id int) error {
	var user User

	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
