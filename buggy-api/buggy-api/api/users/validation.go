package users

import "errors"

func validateNewUserRequest(request *newUserRequest) error {
	if len(request.Username) <= 0 {
		return errors.New("username is required")
	}

	if len(request.Username) > 50 {
		return errors.New("username is too long")
	}

	if len(request.Password) <= 0 {
		return errors.New("password is required")
	}

	if len(request.Password) > 50 {
		return errors.New("password is too long")
	}

	if len(request.FirstName) <= 0 {
		return errors.New("first name is required")
	}

	if len(request.FirstName) > 250 {
		return errors.New("first name is too long")
	}

	if len(request.LastName) <= 0 {
		return errors.New("last name is required")
	}

	if len(request.LastName) > 250 {
		return errors.New("last name is too long")
	}

	return nil
}

func validateUpdateProfileRequest(request *updateProfileRequest) error {
	if len(request.FirstName) <= 0 {
		return errors.New("first name is required")
	}

	if len(request.FirstName) > 250 {
		return errors.New("first name is too long")
	}

	if len(request.LastName) <= 0 {
		return errors.New("last name is required")
	}

	if len(request.LastName) > 250 {
		return errors.New("last name is too long")
	}

	if len(request.Gender) > 250 {
		return errors.New("gender is too long")
	}

	if len(request.Age) > 250 {
		return errors.New("age is too long")
	}

	if len(request.Address) > 500 {
		return errors.New("address is too long")
	}

	if len(request.Phone) > 250 {
		return errors.New("phone is too long")
	}

	if len(request.Hobby) > 250 {
		return errors.New("hobby is too long")
	}

	return nil
}
