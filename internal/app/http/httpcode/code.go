package httpcode

import "net/http"

var (
	Http200 = http.StatusOK
	Http500 = http.StatusInternalServerError
	Http400 = http.StatusBadRequest
	Http404 = http.StatusNotFound
	Http401 = http.StatusUnauthorized
)
