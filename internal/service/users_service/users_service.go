package users_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenTTL = 12 * time.Hour
)

type UsersService struct {
	usersStorage storage.UsersStorage
	cacheStorage storage.CacheStorage
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uint64 `json:"user_id"`
}

func NewUsersService(usersStorage storage.UsersStorage, cacheStorage storage.CacheStorage) *UsersService {
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
	})

	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func (u *UsersService) ParseToken(accessToken string) (uint64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (u *UsersService) GetuserById(ctx context.Context, id uint64) (*model.User, error) {
	key := fmt.Sprintf("user:%d", id)

	// fmt.Println("key: ", key)

	userStr, err := u.cacheStorage.GetFromCache(ctx, key)
	// fmt.Printf("userStr: %s, error: %w", userStr, err)
	if err == redis.Nil {

		user, err := u.usersStorage.GetUserById(id)
		if err != nil {
			return nil, err
		}

		userData, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}

		if err = u.cacheStorage.AddToCache(ctx, key, userData); err != nil {
			return nil, err
		}

		return user, nil
	} else if err != nil {
		return nil, err
	}

	var user = &model.User{}
	err = json.Unmarshal([]byte(userStr), user)
	if err != nil {
		return nil, err
	}

	return user, nil
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
