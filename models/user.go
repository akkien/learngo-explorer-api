package models

import (
	"time"

	"github.com/akkien/explorer-modern/db"
)

// User : user of application
type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

// Session : session
type Session struct {
	ID        int       `json:"name"`
	UUID      string    `json:"uuid"`
	Email     string    `json:"email" validate:"required"`
	UserID    int       `json:"userId" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateSession : Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	// statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	// stmt, err := Db.Prepare(statement)
	// if err != nil {
	// 	return
	// }
	// defer stmt.Close()
	// // use QueryRow to return a row and scan the returned id into the Session struct
	// err = stmt.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

// Session : Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	result := db.DB.Where("user_id = ?", user.ID).Find(&session)
	err = result.Error
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	result := db.DB.Where("uuid = ?", session.UUID).Find(&session)
	err = result.Error
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}

// DeleteByUUID : Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	db.DB.Where("uuid = ?", session.UUID).Delete(&session)
	return
}

// User : Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	result := db.DB.Where("id = ?", session.UserID).First(&user)
	err = result.Error
	return
}

// SessionDeleteAll : Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	result := db.DB.Exec(statement)
	err = result.Error
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	// statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	// stmt, err := Db.Prepare(statement)
	// if err != nil {
	// 	return
	// }
	// defer stmt.Close()

	// // use QueryRow to return a row and scan the returned id into the User struct
	// err = stmt.QueryRow(createUUID(), user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID, &user.UUID, &user.CreatedAt)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	result := db.DB.Delete(&user)
	err = result.Error
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	result := db.DB.Where("id = ?", user.ID).Updates(user)
	err = result.Error
	return
}

// UserDeleteAll : Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	result := db.DB.Exec(statement)
	err = result.Error
	return
}

// Users : Get all users in the database and returns it
func Users() (users []User, err error) {
	result := db.DB.Find(&users)
	err = result.Error
	return
}

// UserByEmail : Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	result := db.DB.Where("email = ?", email)
	err = result.Error
	return
}

// UserByUUID : Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	result := db.DB.Where("uuid = ?", uuid)
	err = result.Error
	return
}
