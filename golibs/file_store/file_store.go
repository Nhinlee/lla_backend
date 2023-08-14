package filestore

import "net/url"

type FileStore interface {
	GenerateResumableUploadURL(objectName string) (*url.URL, error)
	GeneratePublicObjectURL(objectName string) string
}
