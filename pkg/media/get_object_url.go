package media

import (
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v6"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

// GetObjectURL ...
func GetObjectURL(bucketName, objectName string, client *minio.Client, expiry time.Duration) (*url.URL, error) {
	presignedURL, err := client.PresignedGetObject(bucketName, objectName, expiry, url.Values{})
	if err != nil {
		log.Println("Error while getting object URL: ", err.Error())
		return nil, errs.ErrInternal
	}

	return presignedURL, nil
}
