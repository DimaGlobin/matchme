package storage_errors

type StorageError string

func (s StorageError) Error() string {
	return string(s)
}

const (
	userIdAlreadyExists     StorageError = "User with such id already exists"
	phoneNumberAlreadyExist StorageError = "User with such phone number already exists"
	emailAlreadyExists      StorageError = "User with such email already exists"

	likeIdAlreadyExists StorageError = "Like with such id already exists"
	likeAlreadyExists   StorageError = "User has already been liked"

	dislikeIdAlreadyExists StorageError = "Dislike with such id already exists"
	dislikeAlreadyExists   StorageError = "User has already been disliked"

	matchIdAlreadyExists StorageError = "Match with such id already exists"
	matchAlreadyExists   StorageError = "Match already exists"

	fileIdAlreadyExists StorageError = "File with such id already exists"
)

const (
	AlreadyRated StorageError = "User already has been rated"
	SelfRating   StorageError = "Self rating not allowed"
)
