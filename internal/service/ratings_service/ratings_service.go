package ratings_service

import (
	"github.com/DimaGlobin/matchme/internal/mm_errors"
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/cache_storage"
	"github.com/DimaGlobin/matchme/internal/storage/ratings_storage"
	"github.com/DimaGlobin/matchme/internal/storage/users_storage"
)

const (
	like    = "like"
	dislike = "dislike"

	basicRole   = "basic"
	adminRole   = "admin"
	premiumRole = "premium"
)

type RatingsService struct {
	ratingsStorage ratings_storage.RatingsStorage
	usersStorage   users_storage.UsersStorage
	cacheStorage   cache_storage.CacheStorage
}

func NewRatingsService(ratingsStorage ratings_storage.RatingsStorage, usersStorage users_storage.UsersStorage, cacheStorage cache_storage.CacheStorage) *RatingsService {
	return &RatingsService{
		ratingsStorage: ratingsStorage,
		usersStorage:   usersStorage,
		cacheStorage:   cacheStorage,
	}
}

func (r *RatingsService) RecommendUser(userId uint64) (*model.User, error) {
	return r.usersStorage.GetRandomUser(userId)
}

func (r *RatingsService) GetAllLikes(userId uint64) ([]uint64, error) {
	likes, err := r.ratingsStorage.GetAllLikes(userId)
	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (r *RatingsService) AddLike(subjectId, objectId uint64, subjectRole string) (*model.LikeResp, error) {
	id, err := r.ratingsStorage.AddLike(subjectId, objectId)
	if err != nil {
		return nil, err
	}

	var likesLeft *int

	if subjectRole == "basic" {
		likesResp, err := r.cacheStorage.DecLikesCount(subjectId)
		likesLeft = &likesResp

		if err != nil {
			return nil, err
		}

		if likesResp <= 0 {
			return nil, mm_errors.LikesExpired
		}
	}

	exist, err := r.ratingsStorage.CheckLikeExistance(subjectId, objectId)
	if err != nil {
		return nil, err
	}

	if !exist {
		return &model.LikeResp{
			ReactionType: like,
			ReactionId:   id,
			MatchId:      0,
			LikesLeft:    likesLeft,
		}, err
	}

	matchId, err := r.ratingsStorage.AddMatch(subjectId, objectId)
	if err != nil {
		return nil, err
	}

	return &model.LikeResp{
		ReactionType: like,
		ReactionId:   id,
		MatchId:      matchId,
		LikesLeft:    likesLeft,
	}, nil
}

func (r *RatingsService) AddDislike(subjectId, objectId uint64) (*model.DislikeResp, error) {
	id, err := r.ratingsStorage.AddDislike(subjectId, objectId)
	if err != nil {
		return nil, err
	}

	return &model.DislikeResp{
		ReactionType: dislike,
		ReactionId:   id,
	}, err
}
