package entities

type userDAO DataAccessObject

func (dao *userDAO) Insert(u *User) error {
	insertStmt := "INSERT INTO users(username, department, createtime) VALUES(?, ?, ?)"

	stmt, err := dao.Prepare(insertStmt)
	if err != nil {
		return errHandler(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Username, u.Department, u.CreateTime)
	if err != nil {
		return errHandler(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errHandler(err)
	}
	u.UID = int(id)

	return nil
}

func (dao *userDAO) FindAll() []User {
	findAllStmt := "SELECT * FROM users"

	rows, err := dao.Query(findAllStmt)
	errHandler(err)
	defer rows.Close()

	userList := make([]User, 0, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.UID, &u.Username, &u.Department, &u.CreateTime)
		errHandler(err)
		userList = append(userList, u)
	}

	return userList
}

func (dao *userDAO) FindByID(id int) *User {
	findByIDStmt := "SELECT * FROM users WHERE uid = ?"

	stmt, err := dao.Prepare(findByIDStmt)
	errHandler(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	u := User{}
	err = row.Scan(&u.UID, &u.Username, &u.Department, &u.CreateTime)
	errHandler(err)

	return &u
}

func (dao *userDAO) DeleteAll() error {
	deleteAllStmt := "TRUNCATE TABLE users"
	_, err := dao.Exec(deleteAllStmt)
	if err != nil {
		return errHandler(err)
	}
	return nil
}
