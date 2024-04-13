package storage_errors

import (
	"errors"

	"github.com/lib/pq"
)

const (
	postgresUniqueUsersId          = "users_user_id_key"
	postgresUniqueUsersPhoneNumber = "users_phone_number_key"
	postgresUniqueUsersEmail       = "users_email_key"

	postgresUniqueLikeId    = "likes_like_id_key"
	postgresUniqueLike      = "idx_unique_like"
	postgresUniqueDislikeId = "likes_dislike_id_key"
	postgresUniqueDislike   = "idx_unique_dislike"

	postgresUniquematchId = "matches_match_id_key"
	postgresUniqueMatch1  = "idx_unique_match1"
	postgresUniqueMatch2  = "idx_unique_match2"

	postgresUniqueFileId = "files_file_id_key"
)

const (
	uniqueViolation = "unique_violation"
)

func ProcessPostgresError(err error) error {
	var pgErr *pq.Error

	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code.Name() {
	case uniqueViolation:
		switch pgErr.Constraint {
		case postgresUniqueUsersId:
			return userIdAlreadyExists
		case postgresUniqueUsersPhoneNumber:
			return phoneNumberAlreadyExist
		case postgresUniqueUsersEmail:
			return emailAlreadyExists
		case postgresUniqueLikeId:
			return likeIdAlreadyExists
		case postgresUniqueLike:
			return likeAlreadyExists
		case postgresUniqueDislikeId:
			return dislikeIdAlreadyExists
		case postgresUniqueDislike:
			return dislikeAlreadyExists
		case postgresUniquematchId:
			return matchIdAlreadyExists
		case postgresUniqueMatch1, postgresUniqueMatch2:
			return matchAlreadyExists
		case postgresUniqueFileId:
			return fileIdAlreadyExists
		}
	}

	return err
}
