package model

// A Status contains an http status code and its associated error message.
type Status struct {
	HttpStatusCode int
	ErrorMessage   string
}
