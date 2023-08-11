package filestore

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
)

func GetContentType(fileName string) string {
	fileExtension := filepath.Ext(fileName)
	contentType := mime.TypeByExtension(fileExtension)
	if len(contentType) == 0 {
		contentType = "application/octet-stream" //default of google
	}

	return contentType
}

func HeaderToArray(header http.Header) (res []string) {
	for name, values := range header {
		for _, value := range values {
			res = append(res, fmt.Sprintf("%s: %s", name, value))
		}
	}
	return
}
