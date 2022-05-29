package media

import (
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v6"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

// GetAndCopyObject ...
func GetAndCopyObject(bucketName, objectName string, client *minio.Client) error {
	object, err := client.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("Error while getting an object")
		return errs.ErrInternal
	}

	localFile, err := os.Create("./pkg/images/local.jpg")
	if err != nil {
		log.Println("Error while creating a file: ", err)
		return errs.ErrInternal
	}

	if _, err = io.Copy(localFile, object); err != nil {
		log.Println("Error while copying a file", err)
		return errs.ErrInternal
	}

	return nil
}
