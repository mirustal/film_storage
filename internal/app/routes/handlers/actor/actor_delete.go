package actors


// ActorsActorIdDelete - Удаление актёра
func (c *DefaultAPIController) ActorsActorIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	actorIdParam := params["actorId"]
	if actorIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"actorId"}, nil)
		return
	}
	result, err := c.service.ActorsActorIdDelete(r.Context(), actorIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}