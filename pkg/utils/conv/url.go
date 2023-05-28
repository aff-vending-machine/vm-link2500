package conv

import "strings"

func ToURLPath(params ...string) string {
	// Clean up the parameters by trimming whitespaces and removing duplicated "/" characters
	for i, param := range params {
		params[i] = strings.TrimSpace(param)
		params[i] = strings.Trim(params[i], "/")
	}
	// Join the parameters with "/" as a separator
	return strings.Join(params, "/")
}
