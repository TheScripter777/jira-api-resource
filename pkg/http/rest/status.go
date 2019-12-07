package rest

import (
	"errors"
	"fmt"
	"net/http"
)

func HasValidContent(resp *http.Response) bool {
	if resp != nil {
		code := resp.StatusCode
		return codeOK(code) && Is2xx(resp) &&
			code != http.StatusNoContent &&
			resp.Body != nil
	}

	return false
}

func Is4xx(resp *http.Response) (bool, error) {
	code := resp.StatusCode
	if code >= http.StatusBadRequest && code <= 499 {
		return true, errors.New(fmt.Sprintf("Received HTTP%d: %v", code, resp.Body))
	}

	return false, nil
}

func Is5xx(resp *http.Response) (bool, error) {
	code := resp.StatusCode
	if code >= http.StatusInternalServerError && code <= 599 {
		return true, errors.New(fmt.Sprintf("Received HTTP%d: %v", code, resp.Body))
	}

	return false, nil
}

func Is2xx(resp *http.Response) bool {
	code := resp.StatusCode
	return codeOK(code) &&
		code >= http.StatusOK && code <= 299
}

func codeOK(code int) bool {
	return code != 0 && code <= 599
}
