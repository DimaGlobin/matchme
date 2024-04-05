package ratings_storage

import (
	"github.com/jmoiron/sqlx"
)

type RatingsPostgres struct {
	db *sqlx.DB
}

func NewRatingsPostgres(db *sqlx.DB) *RatingsPostgres {
	return &RatingsPostgres{db: db}
}

func (r *RatingsPostgres) AddLike(liking, liked uint64) (uint64, error) {
	var id uint64
	query := "INSERT INTO likes (liking_id, liked_id) values ($1, $2) RETURNING like_id"
	row := r.db.QueryRow(query, liking, liked)

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

func (r *RatingsPostgres) AddDislike(disliking, disliked uint64) (uint64, error) {
	var id uint64
	query := "INSERT INTO dislikes (disliking_id, disliked_id) values ($1, $2) RETURNING dislike_id"
	row := r.db.QueryRow(query, disliking, disliked)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
