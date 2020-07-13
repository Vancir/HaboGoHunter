package pathutil

import (
	"path"
	"strings"
)

func WithSuffix(filepath, suffix string) string {
	// FIXME: validate filepath
	// FIXME: validate suffix
	if strings.HasSuffix(filepath, "/") {
		filepath = strings.TrimSuffix(filepath, "/")
	}
	name := strings.TrimSuffix(filepath, path.Ext(filepath))
	return name + suffix
}
