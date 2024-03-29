package users_storage

import (
	"fmt"
	"strings"

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
	query := "INSERT INTO users (email, phone_number, name, password_hash, sex, age, birth_date, city, description, role, max_age) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING user_id"

	row := u.db.QueryRow(query, user.Email, user.PhoneNumber, user.Name, user.Password, user.Sex, user.Age, user.BirthDate, user.City, user.Description, user.Role, user.MaxAge)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserPostgres) GetUser(email string) (*model.User, error) {
	user := &model.User{}
	// fmt.Printf("email: %s, password: %s", email)
	query := "SELECT * FROM users WHERE email=$1"
	if err := u.db.Get(user, query, email); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgres) GetUserById(id int) (*model.User, error) {
	user := &model.User{}
	query := "SELECT * FROM users where user_id=$1"
	if err := u.db.Get(user, query, id); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgres) UpdateUser(id int, updates model.Updates) error {
	query := "UPDATE users SET"

	count := 1
	var values []interface{}
	for k, v := range updates {
		query += " " + k + fmt.Sprintf(" = $%d,", count)
		switch val := v.(type) {
		case int:
			values = append(values, val)
		case string:
			values = append(values, val)
		case float64:
			values = append(values, int(val))
		default:
			return fmt.Errorf("unsupported type: %T", v)
		}
		count++
	}

	query = strings.TrimSuffix(query, ",")
	values = append(values, id)
	query += fmt.Sprintf(" WHERE user_id = $%d", len(values))

	// fmt.Printf("query: %v\nvalues: %v", query, values)

	_, err := u.db.Exec(query, values...)

	return err
}

func (u *UserPostgres) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE user_id=$1"
	_, err := u.db.Exec(query, id)

	return err
}
