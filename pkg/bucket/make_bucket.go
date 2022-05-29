package bucket

import (
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

// MakeMinioBucket ...
func MakeMinioBucket(minioClient *minio.Client, bucketName, location string) error {
	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Println("This Bucket Already Exists")
			return errs.ErrBucketExists

		} else {
			log.Println("Internal Server Error: ", err.Error())
			return errs.ErrInternal
		}
	}

	return nil
}
