package conv

import "strings"

func JoinURLPathSegments(segments ...string) string {
	var cleanedSegments []string

	// Clean up the segments by trimming whitespaces and removing duplicated "/" characters
	for _, segment := range segments {
		cleanedSegment := strings.TrimSpace(segment)
		cleanedSegment = strings.Trim(cleanedSegment, "/")
		// Add to cleanedSegments only if it's not empty
		if cleanedSegment != "" {
			cleanedSegments = append(cleanedSegments, cleanedSegment)
		}
	}

	// Join the segments with "/" as a separator
	result := strings.Join(cleanedSegments, "/")

	return result
}
