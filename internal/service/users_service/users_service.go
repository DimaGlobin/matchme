package users_service

import (
	"os"
	"time"

	"github.com/DimaGlobin/matchme/internal/mm_errors"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/cache_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL = 12 * time.Hour
)

type UsersService struct {
	usersStorage users_storage.UsersStorage
	cacheStorage cache_storage.CacheStorage
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId   uint64 `json:"user_id"`
	UserRole string `json:"user_role"`
}

func NewUsersService(usersStorage users_storage.UsersStorage, cacheStorage cache_storage.CacheStorage) *UsersService {
	return &UsersService{
		usersStorage: usersStorage,
		cacheStorage: cacheStorage,
	}
}

func (u *UsersService) CreateUser(user *model.User) (uint64, error) {
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hash
	return u.usersStorage.CreateUser(user)
}

func (u *UsersService) GenerateToken(email string, password string) (string, error) {
	user, err := u.usersStorage.GetUser(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func (u *UsersService) ParseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, mm_errors.InvalidSigning
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, mm_errors.InvalidTokenType
	}

	return claims, nil
}

func (u *UsersService) GetuserById(id uint64) (*model.User, error) {
	user, err := u.cacheStorage.GetUserFromCache(id)
	if err == cache_storage.NoValueInCache {
		user, err = u.usersStorage.GetUserById(id)
		if err != nil {
			return nil, err
		}
		if err = u.cacheStorage.AddUserToCache(user); err != nil {
			return nil, err
		}

		return user, nil
	} else if err != nil {
		return nil, err
	}

	return u.usersStorage.GetUserById(id)
}

func (u *UsersService) UpdateUser(id uint64, updates model.Updates) error {
	return u.usersStorage.UpdateUser(id, updates)
}

func (u *UsersService) DeleteUser(id uint64) error {
	return u.usersStorage.DeleteUser(id)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
