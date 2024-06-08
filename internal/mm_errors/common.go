package mm_errors

const (
	DecodeError      MmError = "Unable to decode request body"
	JwtCreationError MmError = "Unable to create jwt token"
)
