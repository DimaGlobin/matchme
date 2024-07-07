package ratings_storage

import "github.com/DimaGlobin/matchme/internal/model"

type RatingsStorage interface {
	AddLike(like *model.Like) (uint64, error)
	AddDislike(dislike *model.Dislike) (uint64, error)
	GetAllLikes(userId uint64) ([]uint64, error)
	AddMatch(*model.Match) (uint64, error)
	CheckLikeExistance(likingId, likedId uint64) (bool, error)
	CheckDislikeExistance(dislikingId, dislikedId uint64) (bool, error)
}

var _ RatingsStorage = (*RatingsPostgres)(nil)
