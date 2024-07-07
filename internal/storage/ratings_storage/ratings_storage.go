package ratings_storage

import (
	"github.com/DimaGlobin/matchme/internal/model"
	"github.com/DimaGlobin/matchme/internal/storage/storage_errors"
	"gorm.io/gorm"
)

type RatingsPostgres struct {
	db *gorm.DB
}

func NewRatingsPostgres(db *gorm.DB) *RatingsPostgres {
	return &RatingsPostgres{db: db}
}

func (r *RatingsPostgres) AddLike(like *model.Like) (uint64, error) {
	if err := r.db.Create(like).Error; err != nil {
		return 0, storage_errors.ProcessPostgresError(err)
	}
	return like.LikedId, nil
}

func (r *RatingsPostgres) GetAllLikes(userId uint64) ([]uint64, error) {
	var ids []uint64
	query := `
	SELECT liking_id 
	FROM likes 
	WHERE liked_id = ? 
	AND liking_id NOT IN (
		SELECT disliked_id 
		FROM dislikes 
		WHERE disliking_id = ?
	)
	`

	rows, err := r.db.Raw(query, userId, userId).Rows()
	if err != nil {
		return nil, storage_errors.ProcessPostgresError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, storage_errors.ProcessPostgresError(err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, storage_errors.ProcessPostgresError(err)
	}

	return ids, nil
}

func (r *RatingsPostgres) AddDislike(dislike *model.Dislike) (uint64, error) {
	if err := r.db.Create(dislike).Error; err != nil {
		return 0, storage_errors.ProcessPostgresError(err)
	}
	return dislike.Id, nil
}

func (r *RatingsPostgres) AddMatch(match *model.Match) (uint64, error) {
	if err := r.db.Create(match).Error; err != nil {
		return 0, storage_errors.ProcessPostgresError(err)
	}
	return match.MatchId, nil
}

func (r *RatingsPostgres) CheckLikeExistance(likingId, likedId uint64) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Like{}).
		Where("liking_id = ? AND liked_id = ?", likedId, likingId).
		Count(&count).Error; err != nil {
		return false, storage_errors.ProcessPostgresError(err)
	}

	return count > 0, nil
}

func (r *RatingsPostgres) CheckDislikeExistance(dislikingId, dislikedId uint64) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Dislike{}).
		Where("disliking_id = ? AND disliked_id = ?", dislikedId, dislikingId).
		Count(&count).Error; err != nil {
		return false, storage_errors.ProcessPostgresError(err)
	}

	return count > 0, nil
}
