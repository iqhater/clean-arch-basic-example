package controller

import (
	"context"
	"net/http"
	"net/url"
)

// ValidateRequest middleware handler check and validate user requset
func ValidateRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// check on correct request method
		if req.Method != http.MethodGet {
			http.Error(w, "Wrong method used! Only GET method allowed.", http.StatusMethodNotAllowed)
			return
		}

		// parse encoded query params to map
		params, err := url.ParseQuery(req.URL.Query().Encode())
		if err != nil {
			http.Error(w, "Invalid query parameters!", http.StatusBadRequest)
			return
		}

		// check if "name" word is exists
		if _, ok := params[string(contextNameKey)]; !ok {
			http.Error(w, "'name' parameter does not exist or bad value!", http.StatusBadRequest)
			return
		}

		// parse name and check if it is empty
		name := params.Get(string(contextNameKey))
		if name == "" {
			http.Error(w, "'name' value is empty!", http.StatusBadRequest)
			return
		}

		// add valid name value to context
		ctx := context.WithValue(req.Context(), contextNameKey, name)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
