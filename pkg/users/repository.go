package users

type userRepository struct {
	users map[int]*User
	seq   int
}

func NewUserRepository() UserRepository {
	users := map[int]*User{}
	seq := 1
	return &userRepository{
		users: users,
		seq:   seq,
	}
}

func (u *userRepository) GetAll() ([]*User, error) {
	users := []*User{}
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetById(id int) (*User, error) {
	users := u.users
	return users[id], nil
}

func (u *userRepository) Create(usr *User) (*User, error) {
	users := u.users
	usr.ID = u.seq
	users[usr.ID] = usr
	u.seq++
	return usr, nil
}

func (u *userRepository) Update(id int, usr *User) (*User, error) {
	users := u.users
	users[id].Name = usr.Name
	return users[id], nil
}

func (u *userRepository) Delete(id int) error {
	users := u.users
	delete(users, id)
	return nil
}
