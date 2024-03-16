package models

type Actor struct {
	// Уникальный идентификатор актёра
	Id string `json:"id,omitempty"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthDate"`
	films []string
}

type Movie struct {
	// Уникальный идентификатор фильма
	Id string `json:"id,omitempty"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
	Rating float32 `json:"rating"`
	Actors []string
}


type User struct {
	Id int64 `json:"id,omitempty"`
	Role string `json:"role"`
}