package model

type Updates map[string]interface{}

func (u Updates) Valid() bool {

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
