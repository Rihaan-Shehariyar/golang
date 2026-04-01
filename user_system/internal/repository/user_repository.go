package repository

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int32
	Name  string
	Email string
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(name, email string) (*User, error) {
	var id int32

	err := r.DB.QueryRow(
		"INSERT INTO users(name,email) VALUES($1, $2) RETURNING id",
		name, email,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &User{ID: id, Name: name, Email: email}, nil

}

func (r *UserRepository) GetUser(id int32) (*User, error) {

	row := r.DB.QueryRow("SELECT id,name,email FROM users WHERE id = &1", id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
func (r *UserRepository) ListUser() ([]*User, error) {
	rows, err := r.DB.Query("SELECT id,name,email from users")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
