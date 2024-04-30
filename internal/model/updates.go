package model

type Updates struct {
	Updates map[string]interface{} `json:"updates" swaggertype:"object,string" example:"email:adb@wda.com, name:Sara"`
}
type UpdateError string

func (u UpdateError) Error() string {
	return string(u)
}

const (
	emptyBody         UpdateError = "Empty updates"
	passwordChange    UpdateError = "Cannot change password"
	birthDateChanging UpdateError = "Cannot change burth date"
)

func (u Updates) Valid() error {

	if len(u.Updates) <= 0 {
		return emptyBody
	}

	if _, ok := u.Updates["password"]; ok {
		return passwordChange
	}

	if _, ok := u.Updates["birth_date"]; ok {
		return birthDateChanging
	}

	return nil
}
