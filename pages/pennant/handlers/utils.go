package handlers

import (
	"fmt"
	"strings"
)

/* Substitute values into push URL if it contains parameters */
func SubstitutePushUrlParams(pushUrl string, substituteValues map[string]string) string {
	var substitutedUrl string
	params := strings.Split(pushUrl, "/")
	for _, param := range params {
		if param[0] == ':' {
			substitutedUrl = fmt.Sprintf("%s/%s", substitutedUrl, substituteValues[param[1:]])
		} else {
			substitutedUrl = fmt.Sprintf("%s/%s", substitutedUrl, param)
		}
	}
	return substitutedUrl
}