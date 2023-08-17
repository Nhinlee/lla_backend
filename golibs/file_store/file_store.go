package filestore

import (
	"context"
	"net/url"
)

type FileStore interface {
	GenerateResumableUploadURL(objectName string) (*url.URL, error)
	GeneratePublicObjectURL(objectName string) string
	MoveObject(ctx context.Context, srcObjectName string, destObjectName string) error
}
