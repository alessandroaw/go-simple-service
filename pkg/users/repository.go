package user

import (
	"github.com/therealsandro/go-simple-service/pkg/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	db.AutoMigrate(&domain.User{})

	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetById(id int) (domain.User, error) {
	var user domain.User
	if err := u.db.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *userRepository) Create(usr *domain.User) (*domain.User, error) {
	if err := u.db.Create(usr).Error; err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userRepository) Update(id int, newUsr *domain.User) (*domain.User, error) {
	var user domain.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	u.db.Model(&user).Updates(newUsr)

	return &user, nil

}

func (u *userRepository) Delete(id int) error {
	var user domain.User

	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
