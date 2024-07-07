package users_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/storage_errors"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUsersPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) CreateUser(user *model.User) (uint64, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, storage_errors.ProcessPostgresError(result.Error)
	}

	return user.Id, nil
}

func (u *UserPostgres) GetRandomUser(userId uint64) (*model.User, error) {
    user, err := u.GetUserById(userId)
    if err != nil {
        return nil, err
    }

    var recUserId uint64
    query := `
    SELECT user_id
    FROM users
    WHERE user_id != ?
    AND user_id NOT IN (
        SELECT liked_id FROM likes WHERE liking_id = ?
        UNION
        SELECT disliked_id FROM dislikes WHERE disliking_id = ?
    )
    AND user_id NOT IN (
        SELECT disliking_id FROM dislikes WHERE disliked_id = ?
    )
    AND sex != ?
    ORDER BY RANDOM()
    LIMIT 1;
    `

    result := u.db.Raw(query, userId, userId, userId, userId, user.Sex).Scan(&recUserId)
    if result.Error != nil {
        return nil, storage_errors.ProcessPostgresError(result.Error)
    }

    return u.GetUserById(recUserId)
}

func (u *UserPostgres) GetUser(email string) (*model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, storage_errors.ProcessPostgresError(result.Error)
	}
	return &user, nil
}

func (u *UserPostgres) GetUserById(id uint64) (*model.User, error) {
	var user model.User
	result := u.db.First(&user, id)
	if result.Error != nil {
		return nil, storage_errors.ProcessPostgresError(result.Error)
	}
	return &user, nil
}

func (u *UserPostgres) UpdateUser(id uint64, updates model.Updates) error {
	result := u.db.Model(&model.User{}).Where("user_id = ?", id).Updates(updates.Updates)
	if result.Error != nil {
		return storage_errors.ProcessPostgresError(result.Error)
	}
	return nil
}

func (u *UserPostgres) DeleteUser(id uint64) error {
	result := u.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return storage_errors.ProcessPostgresError(result.Error)
	}
	return nil
}
