package helper

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func NewClientUploader(ctx context.Context) (*ClientUploader, error) {
	var client *storage.Client
	var err error

	if strings.ToLower(os.Getenv("ENVIRONMENT")) == "production" {
		client, err = storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile("credentials.json"))
		if err != nil {
			return nil, err
		}
	}

	uploader := &ClientUploader{
		cl:         client,
		bucketName: os.Getenv("STORAGE_BUCKET_ARTICLE"),
		projectID:  os.Getenv("PROJECT_ID"),
		uploadPath: os.Getenv("STORAGE_UPLOAD_PATH"),
	}

	return uploader, nil
}

func (c *ClientUploader) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error) {

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + "/" + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	fileURL := "https://storage.googleapis.com/" + c.bucketName + "/" + c.uploadPath + "/" + fileName
	return fileURL, nil
}

func (c *ClientUploader) DeleteFile(ctx context.Context, fileName string) error {
	obj := c.cl.Bucket(c.bucketName).Object(c.uploadPath + "/" + fileName)
	return obj.Delete(ctx)
}
