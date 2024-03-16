package validators

import (
	"errors"
	"film_storage/internal/models"
	"film_storage/pkg/utils"
)




func AssertActorRequired(obj models.Actor) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"gender": obj.Gender,
		"birthDate": obj.BirthDate,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertActorConstraints checks if the values respects the defined constraints
func AssertActorConstraints(obj models.Actor) error {
	return nil
}


// AssertMovieRequired checks if the required fields are not zero-ed
func AssertMovieRequired(obj models.Movie) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"description": obj.Description,
		"releaseDate": obj.ReleaseDate,
		"rating": obj.Rating,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMovieConstraints checks if the values respects the defined constraints
func AssertMovieConstraints(obj models.Movie) error {
	if obj.Rating < 0 {
		return &utils.ParsingError{Err: errors.New(utils.ErrMsgMinValueConstraint)}
	}
	if obj.Rating > 10 {
		return &utils.ParsingError{Err: errors.New(utils.ErrMsgMaxValueConstraint)}
	}
	return nil
}


// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj models.User) error {
	elements := map[string]interface{}{
		"role": obj.Role,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj models.User) error {
	return nil
}