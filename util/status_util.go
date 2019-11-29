package util

import (
	"net/http"

	. "backend/model"
)

// StatusOK operates as a constant named StatusOK.
func StatusOK() Status {
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   ""}
}

// IsStatusOK returns true if the given status is StatusOK.
func IsStatusOK(status Status) bool {
	return status.HttpStatusCode == http.StatusOK
}
