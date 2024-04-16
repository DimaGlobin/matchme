package ratings_service

import "github.com/DimaGlobin/matchme/internal/model"

type RatingsError string

func (r RatingsError) Error() string {
	return string(r)
}

const (
	unsupReaction RatingsError = "Unsuported reaction"
)

type RatingsServiceInt interface {
	RecommendUser(userId uint64) (*model.User, error)
	AddReaction(reaction string, subjectId, objectId uint64) (uint64, uint64, error)
	GetAllLikes(userId uint64) ([]uint64, error)
}

var _ RatingsServiceInt = (*RatingsService)(nil)
