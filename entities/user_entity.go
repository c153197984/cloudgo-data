package entities

import "time"

// A User entity.
//
// UID - The id of the user.
// Username - The name of the user.
// Department - The department name of the user.
// CreateTime - The time when the user is created.
type User struct {
	UID        int        `xorm:"'uid' INT(10) PK NOT NULL AUTOINCR"`
	Username   string     `xorm:"'username' VARCHAR(64) NULL DEFAULT NULL"`
	Department string     `xorm:"'department' VARCHAR(64) NULL DEFAULT NULL"`
	CreateTime *time.Time `xorm:"'createtime' DATE NULL DEFAULT NULL"`
}

// NewUser returns a new instance of a User.
func NewUser(username string, department string) *User {
	u := User{}
	if len(username) == 0 {
		panic("Username should not be empty.")
	}
	if len(department) == 0 {
		panic("Department should not be empty.")
	}
	t := time.Now()
	u.Username = username
	u.Department = department
	u.CreateTime = &t

	return &u
}
