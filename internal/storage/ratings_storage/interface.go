package ratings_storage

type RatingsStorage interface {
	AddLike(likingId, likedId uint64) (uint64, error)
	AddDislike(dislikingId, dislikedId uint64) (uint64, error)
	GetAllLikes(userId uint64) ([]uint64, error)
	AddMatch(userId1, userId2 uint64) (uint64, error)
	CheckLikeExistance(likingId, likedId uint64) (bool, error)
}

var _ RatingsStorage = (*RatingsPostgres)(nil)
