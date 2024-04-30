package ratings_service

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/ratings_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
)

const (
	like    = "like"
	dislike = "dislike"
)

type RatingsService struct {
	ratingsStorage ratings_storage.RatingsStorage
	usersStorage   users_storage.UsersStorage
}

func NewRatingsService(ratingsStorage ratings_storage.RatingsStorage, usersStorage users_storage.UsersStorage) *RatingsService {
	return &RatingsService{
		ratingsStorage: ratingsStorage,
		usersStorage:   usersStorage,
	}
}

func (r *RatingsService) RecommendUser(userId uint64) (*model.User, error) {
	return r.usersStorage.GetRandomUser(userId)
}

func (r *RatingsService) AddReaction(reaction string, subjectId, objectId uint64) (uint64, uint64, error) {
	if reaction == like {
		id, err := r.ratingsStorage.AddLike(subjectId, objectId)
		if err != nil {
			return 0, 0, err
		}

		exist, err := r.ratingsStorage.CheckLikeExistance(subjectId, objectId)
		if err != nil {
			return 0, 0, err
		}

		if !exist {
			return id, 0, nil
		}

		matchId, err := r.ratingsStorage.AddMatch(subjectId, objectId)
		if err != nil {
			return 0, 0, err
		}

		return id, matchId, nil

	} else if reaction == dislike {
		id, err := r.ratingsStorage.AddDislike(subjectId, objectId)
		if err != nil {
			return 0, 0, err
		}

		return id, 0, nil
	}

	return 0, 0, unsupReaction
}

func (r *RatingsService) GetAllLikes(userId uint64) ([]uint64, error) {
	likes, err := r.ratingsStorage.GetAllLikes(userId)
	if err != nil {
		return nil, err
	}

	return likes, nil
}
