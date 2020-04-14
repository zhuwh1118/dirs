package stringx

import (
	"fmt"
	"strings"
)

func TrimPrefixSlash(path string) string {
	if strings.HasPrefix(path, "/") {
		for {
			path = strings.TrimPrefix(path, "/")
			if !strings.HasPrefix(path, "/") {
				break
			}
		}
	}
	path = fmt.Sprintf("/%s", path)
	return path
}
