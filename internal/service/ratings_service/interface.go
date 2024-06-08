package ratings_service

import "github.com/DimaGlobin/matchme/internal/model"

type RatingsError string

func (r RatingsError) Error() string {
	return string(r)
}

type RatingsServiceInt interface {
	RecommendUser(userId uint64) (*model.User, error)
	AddLike(subjectId, objectId uint64, subjectRole string) (*model.LikeResp, error)
	AddDislike(subjectId, objectId uint64) (*model.DislikeResp, error)
	GetAllLikes(userId uint64) ([]uint64, error)
}

var _ RatingsServiceInt = (*RatingsService)(nil)
