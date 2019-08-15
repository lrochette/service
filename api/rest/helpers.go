package rest

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

// APIError is the returned error
type APIError struct {
	Message string `json:"message"`
}

// decodeRequest writes a REST request body to the given requestObject.
// The function returns true if the decode was successful, false otherwise.
func decodeRequest(request *restful.Request, response *restful.Response, requestObject interface{}) bool {
	if err := request.ReadEntity(&requestObject); err != nil {
		encodeErrorWithStatus(response, err, http.StatusBadRequest)
		return false
	}
	return true
}

// encodeErrorWithStatus writes the given error and status code to a REST response.
func encodeErrorWithStatus(response *restful.Response, err error, status int) {
	apierr := &APIError{
		Message: err.Error(),
	}

	response.WriteHeaderAndEntity(status, apierr)
}
