package user

import (
	"film_storage/pkg/utils"
	"net/http"
)


func TokenRefreshPost(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.TokenRefreshPost(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}