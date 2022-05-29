package media

import (
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

// PutObject ...
func PutObject(client *minio.Client, bucketName, objectName, filePath, contentType string, options minio.PutObjectOptions) error {
	// Upload the image file with FPutObject
	_, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return errs.ErrInternal
	}

	return nil
}
