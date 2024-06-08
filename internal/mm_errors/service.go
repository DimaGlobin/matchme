package mm_errors

const (
	InvalidSigning   MmError = "Invalid signing method"
	InvalidTokenType MmError = "Token claims are not of type *tokenClaims"

	LikesExpired MmError = "Likes count expired"
)
