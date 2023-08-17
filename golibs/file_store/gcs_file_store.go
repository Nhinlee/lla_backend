package filestore

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	bucketName = "llapp-bucket"
	saKeyEnv   = "GOOGLE_SERVICE_ACCOUNT_KEY"
	saEmailEnv = "GOOGLE_SERVICE_ACCOUNT"
)

type GCSFileStore struct {
	client              *storage.Client
	httpClient          *http.Client
	serviceAccountEmail string
}

func NewGCSFileStore() (*GCSFileStore, error) {
	ctx := context.Background()

	saEmail := os.Getenv(saEmailEnv)

	serviceAccountKeyEncoded := os.Getenv(saKeyEnv)
	serviceAccountKeyStr, err := base64.StdEncoding.DecodeString(serviceAccountKeyEncoded)
	if err != nil {
		return nil, err
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(serviceAccountKeyStr)))
	if err != nil {
		return nil, err
	}

	return &GCSFileStore{
		client: client,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		serviceAccountEmail: saEmail,
	}, nil
}

func (s *GCSFileStore) GenerateResumableUploadURL(objectName string) (*url.URL, error) {
	contentType := GetContentType(objectName)

	opts := &storage.SignedURLOptions{
		GoogleAccessID: s.serviceAccountEmail,
		Method:         http.MethodPost,
		Scheme:         storage.SigningSchemeV4,
		ContentType:    contentType,
		Expires:        time.Now().Add(10 * time.Minute),
		QueryParameters: url.Values{
			"name": {objectName},
		},
		Headers: HeaderToArray(http.Header{
			"x-goog-resumable":      {"start"},
			"X-Upload-Content-Type": {contentType},
			"uploadType":            {"resumable"},
		}),
	}

	presignedURL, err := s.client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return nil, fmt.Errorf("cannot generate presigned URL: %w", err)
	}

	return s.initResumableUploadSession(http.MethodPost, presignedURL, contentType, "*")
}

func (s *GCSFileStore) initResumableUploadSession(method, signUrl, contentType, allowOrigin string) (*url.URL, error) {
	req, err := http.NewRequest(method, signUrl, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("could not init resumable upload session: %v", err)
	}
	req.Header.Set("X-Upload-Content-Type", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-goog-resumable", "start")
	req.Header.Set("Origin", allowOrigin)
	req.Header.Set("uploadType", "resumable")

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not init resumable upload session: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("could not init resumable upload session: status code:%d, header:%s, body:%s", res.StatusCode, res.Header, bodyString)
	}
	sessionURI := res.Header.Get("Location")

	return url.Parse(sessionURI)
}

func (s *GCSFileStore) GeneratePublicObjectURL(objectName string) string {
	dir, filename := filepath.Dir(objectName), filepath.Base(objectName)
	if dir == "." {
		filename = url.PathEscape(filename)
	} else {
		filename = fmt.Sprintf("%s/%s", dir, url.PathEscape(filename))
	}
	return fmt.Sprintf("%s/%s/%s", "https://storage.googleapis.com", bucketName, filename)
}

func (g *GCSFileStore) MoveObject(ctx context.Context, srcObjectName, destObjetName string) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}

	src := client.Bucket(bucketName).Object(srcObjectName)
	dst := client.Bucket(bucketName).Object(destObjetName)

	if _, err := src.Attrs(ctx); err != nil {
		return fmt.Errorf("Object(%s/%s).Attrs: %v", bucketName, srcObjectName, err)
	}

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return fmt.Errorf("Object(%s).CopierFrom(%s).Run: %v", bucketName, srcObjectName, err)
	}

	if err := src.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%s/%s).Delete: %v", bucketName, srcObjectName, err)
	}
	return nil
}
