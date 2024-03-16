package user

import (
	"encoding/json"
	"film_storage/pkg/utils"
)

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	query, err := utils.ParseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	registerPostRequestParam := RegisterPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&registerPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRegisterPostRequestRequired(registerPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRegisterPostRequestConstraints(registerPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var adminModeParam string
	if query.Has("adminMode") {
		param := query.Get("adminMode")

		adminModeParam = param
	} else {
	}
	result, err := c.service.RegisterPost(r.Context(), registerPostRequestParam, adminModeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}