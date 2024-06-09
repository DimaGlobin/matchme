package mm_errors

type MmError string

func (m MmError) Error() string {
	return string(m)
}

func NewMmError(msg string) MmError {
	return MmError(msg)
}
