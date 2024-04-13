package model

type Updates map[string]interface{}

type UpdateError string

func (u UpdateError) Error() string {
	return string(u)
}

const (
	emptyBody UpdateError = "Empty updates"
	passwordChange UpdateError = "Cannot change password"
	birthDateChanging UpdateError = "Cannot change burth date"
)

func (u Updates) Valid() error {

	if len(u) <= 0 {
		return false
	}

	if _, ok := u["password"]; ok {
		return false
	}

	if _, ok := u["birth_date"]; ok {
		return false
	}

	return true
}
