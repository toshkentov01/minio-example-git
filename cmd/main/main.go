package main

import (
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v6"
	"github.com/sardortoshkentov/minio-example/config"
	"github.com/sardortoshkentov/minio-example/pkg/bucket"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
	getobjectname "github.com/sardortoshkentov/minio-example/pkg/get_object_name"
	"github.com/sardortoshkentov/minio-example/pkg/media"
	uuidgenerator "github.com/sardortoshkentov/minio-example/pkg/uuid_generator"
)

func main() {
	var (
		cfg                = config.Get()
		filePath    string = "./pkg/images/ubuntu2.jpg"
		contentType string = "image/jpg"
	)

	// Comment this if you don't use put function
	objectName, err := uuidgenerator.GenerateUUID()

	// Initialize minio client object.
	minioClient, err := minio.New(cfg.EndPoint, cfg.AccessID, cfg.SecretKey, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Comment or not call this function after creating a bucket
	err = bucket.MakeMinioBucket(minioClient, cfg.BucketName, cfg.Location)
	if err != nil {
		if err == errs.ErrBucketExists {
			log.Println("Be careful, this bucket already exists")
			
		} else {
			log.Fatal("Something went wrong, error: ", err)
			return
		}
	}

	// Comment or not call this function after putting the object (Do not put same object twice)
	err = media.PutObject(minioClient, cfg.BucketName, objectName.String(), filePath, contentType, minio.PutObjectOptions{})
	if err != nil {
		log.Println("Something went wrong: ", err.Error())
		return
	}

	// Comment or not call this function if you don't use it
	createdObjectName, err := getobjectname.GetObjectName()
	if err != nil {
		log.Println("Something went wrong: ", err.Error())
		return
	}

	// Comment or not call this function if you don't use it
	err = media.GetAndCopyObject(cfg.BucketName, createdObjectName, minioClient)
	if err != nil {
		log.Println("Something went wrong: ", err.Error())
		return
	}

	expiry := time.Second * 24 * 60 * 60 // 1 day.

	objectURL, err := media.GetObjectURL(cfg.BucketName, createdObjectName, minioClient, expiry)
	if err != nil {
		log.Println("Something went wrong: ", err.Error())
		return
	}

	log.Println(objectURL)
	log.Println("Success")
}
