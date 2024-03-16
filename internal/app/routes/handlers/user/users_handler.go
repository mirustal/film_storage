package user

import (
	"film_storage/pkg/utils"
)




type RegisterDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin bool `json:"admin,omitempty"`
}

// AssertRegisterPostRequestRequired checks if the required fields are not zero-ed
func AssertRegisterPostRequestRequired(obj RegisterDTO) error {
	elements := map[string]interface{}{
		"username": obj.Username,
		"password": obj.Password,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRegisterPostRequestConstraints checks if the values respects the defined constraints
func AssertRegisterPostRequestConstraints(obj RegisterDTO) error {
	return nil
}

