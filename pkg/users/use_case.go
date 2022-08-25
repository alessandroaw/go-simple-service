package users

// userUseCase is the object that will be used to implement the UserUseCase interface
type userUseCase struct {
	userRepo UserRepository
}

/*
NewUserUseCase will initialize the userUseCase
the useeUseCase itself does not need to be exported because it is
not intended to be used outside of the package
however it must comply with the UserUseCase interface
*/
func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) GetAll() ([]*User, error) {
	users, err := u.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUseCase) GetById(id int) (*User, error) {
	usr, err := u.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userUseCase) Create(usr *User) (*User, error) {
	usr, err := u.userRepo.Create(usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userUseCase) Update(id int, usr *User) (*User, error) {
	usr, err := u.userRepo.Update(id, usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *userUseCase) Delete(id int) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
