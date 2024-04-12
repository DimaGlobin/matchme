package ratings_storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RatingsPostgres struct {
	db *sqlx.DB
}

func NewRatingsPostgres(db *sqlx.DB) *RatingsPostgres {
	return &RatingsPostgres{db: db}
}

func (r *RatingsPostgres) AddLike(likingId, likedId uint64) (uint64, error) {
	var id uint64

	if likingId == likedId {
		return 0, fmt.Errorf("Not allowed to like yourself")
	}

	query := `
	INSERT INTO likes (liking_id, liked_id)
	VALUES ($1, $2)
	RETURNING like_id;
	`
	row := r.db.QueryRow(query, likingId, likedId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RatingsPostgres) GetAllLikes(userId uint64) ([]uint64, error) {
	var ids []uint64
	query := "SELECT liking_id FROM likes WHERE liked_id=$1"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *RatingsPostgres) AddDislike(dislikingId, dislikedId uint64) (uint64, error) {
	var id uint64

	if dislikingId == dislikedId {
		return 0, fmt.Errorf("Not allowed to like yourself")
	}
	query := "INSERT INTO dislikes (disliking_id, disliked_id) values ($1, $2) RETURNING dislike_id"
	row := r.db.QueryRow(query, dislikingId, dislikedId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RatingsPostgres) AddMatch(userId1, userId2 uint64) (uint64, error) {
	var id uint64
	query := "INSERT INTO matches (user1, user2) values ($1, $2) RETURNING match_id"
	row := r.db.QueryRow(query, userId1, userId2)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RatingsPostgres) CheckLikeExistance(likingId, likedId uint64) (bool, error) {
	var count uint
	query := "SELECT COUNT(*) FROM likes where liking_id=$1 AND liked_id=$2"
	if err := r.db.Get(&count, query, likedId, likingId); err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
