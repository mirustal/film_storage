package handler

import (
	"film_storage/pkg/utils"
	"net/http"
)


type ApiRoute interface { 
	ActorsActorIdDelete(http.ResponseWriter, *http.Request)
	ActorsActorIdMoviesGet(http.ResponseWriter, *http.Request)
	ActorsActorIdPatch(http.ResponseWriter, *http.Request)
	ActorsPost(http.ResponseWriter, *http.Request)
	MoviesGet(http.ResponseWriter, *http.Request)
	MoviesMovieIdDelete(http.ResponseWriter, *http.Request)
	MoviesMovieIdPatch(http.ResponseWriter, *http.Request)
	MoviesPost(http.ResponseWriter, *http.Request)
	RegisterPost(http.ResponseWriter, *http.Request)
	TokenGet(http.ResponseWriter, *http.Request)
	TokenRefreshPost(http.ResponseWriter, *http.Request)
}
