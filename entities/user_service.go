package entities

// A UserServiceProvider provides a series of
// operations to Entity User.
type UserServiceProvider struct{}

// UserService is an instance of UserAtomicService.
var UserService = UserServiceProvider{}

// Insert is an atomic operation that inserts a new
// user to the database.
func (*UserServiceProvider) Insert(u *User) error {
	tx, err := db.Begin()
	if err != nil {
		return errHandler(err)
	}

	dao := userDAO{tx}
	err = dao.Insert(u)
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll returns all users in the database.
func (*UserServiceProvider) FindAll() []User {
	dao := userDAO{db}
	return dao.FindAll()
}

// FindByID returns a User with a specific ID.
func (*UserServiceProvider) FindByID(id int) *User {
	dao := userDAO{db}
	return dao.FindByID(id)
}

// DeleteAll removes all users in the table.
func (*UserServiceProvider) DeleteAll() error {
	dao := userDAO{db}
	err := dao.DeleteAll()
	if err != nil {
		return errHandler(err)
	}
	return nil
}
