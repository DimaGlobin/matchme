package mm_errors

// jwt --
const (
	DecodeError      MmError = "Unable to decode request body"
	JwtCreationError MmError = "Unable to create jwt token"
)
