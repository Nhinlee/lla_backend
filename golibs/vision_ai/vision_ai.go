package vision_ai

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"lla/golibs"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

type VisionAI struct {
}

func NewVisionAI() *VisionAI {
	serviceAccountKeyEncoded := os.Getenv(golibs.SAKeyEnv)
	serviceAccountKeyStr, err := base64.StdEncoding.DecodeString(serviceAccountKeyEncoded)
	if err != nil {
		panic(err)
	}

	// Write the service account key to a JSON file
	// TODO: temporary solution, need to save the service account key to secret manager
	err = os.WriteFile(".sa.json", serviceAccountKeyStr, 0644)
	if err != nil {
		panic(err)
	}

	os.Setenv(golibs.GoogleCredentials, ".sa.json")

	return &VisionAI{}
}

func (v *VisionAI) DetectLabelsFromImageURI(w io.Writer, fileName string) ([]string, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	image := vision.NewImageFromURI(fmt.Sprintf("gs://%s/%s", golibs.BucketName, fileName))
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	var labels []string
	for _, annotation := range annotations {
		labels = append(labels, annotation.Description)
	}

	return labels, nil
}
