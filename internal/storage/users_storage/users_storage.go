package users_storage

import (
	"fmt"
	"strings"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/storage_errors"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) CreateUser(user *model.User) (uint64, error) {
	var id uint64
	query := "INSERT INTO users (email, phone_number, name, password_hash, sex, age, birth_date, city, description, max_age) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id"

	row := u.db.QueryRow(query, user.Email, user.PhoneNumber, user.Name, user.Password, user.Sex, user.Age, user.BirthDate, user.City, user.Description, user.MaxAge)
	if err := row.Scan(&id); err != nil {
		return 0, storage_errors.ProcessPostgresError(err)
	}

	return id, nil
}

func (u *UserPostgres) GetRandomUser(userId uint64) (*model.User, error) {

	var recUserId uint64
	user, err := u.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT user_id
	FROM users
	WHERE user_id != $1
	AND user_id NOT IN (
		SELECT liked_id FROM likes WHERE liking_id = $1
		UNION
		SELECT disliked_id FROM dislikes WHERE disliking_id = $1
	)
	AND user_id NOT IN (
		SELECT disliking_id FROM dislikes WHERE disliked_id = $1
	)
	AND sex != $2
	ORDER BY RANDOM()
	LIMIT 1;
	`
	//TODO: описать тест кейсы для такого функицонала

	if err := u.db.Get(&recUserId, query, userId, user.Sex); err != nil {
		return nil, err
	}

	return u.GetUserById(recUserId)
}

func (u *UserPostgres) GetUser(email string) (*model.User, error) {
	user := &model.User{}
	// fmt.Printf("email: %s, password: %s", email)
	query := "SELECT * FROM users WHERE email=$1"
	if err := u.db.Get(user, query, email); err != nil {
		return nil, storage_errors.ProcessPostgresError(err)
	}

	return user, nil
}

func (u *UserPostgres) GetUserById(id uint64) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users where user_id=$1"
	if err := u.db.Get(user, query, id); err != nil {
		return nil, storage_errors.ProcessPostgresError(err)
	}

	return user, nil
}

func (u *UserPostgres) UpdateUser(id uint64, updates model.Updates) error {
	query := "UPDATE users SET"

	count := 1
	var values []interface{}
	for k, v := range updates.Updates {
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

func (u *UserPostgres) DeleteUser(id uint64) error {
	query := "DELETE FROM users WHERE user_id=$1"
	_, err := u.db.Exec(query, id)

	return err
}
