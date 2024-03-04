package users_storage

import (
	"fmt"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(user *model.User) (int, error) {
	var id int
	query := "INSERT INTO users (email, phone_number, name, password_hash, sex, age, birth_date, city, description, role, max_age) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"

	row := u.db.QueryRow(query, user.Email, user.PhoneNumber, user.Name, user.Password, user.Sex, user.Age, user.BirthDate, user.City, user.Description, user.Role, user.MaxAge)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserPostgres) GetUser(email string) (*model.User, error) {
	user := &model.User{}
	fmt.Printf("email: %s, password: %s", email)
	query := "SELECT * FROM users WHERE email=$1"
	err := u.db.Get(user, query, email)

	return user, err
}
