package errs

import "errors"

var (
	// ErrBucketExists ...
	ErrBucketExists = errors.New("this bucket already exists")

	// ErrWhileConnecting ...
	ErrWhileConnecting = errors.New("error while connecting to minio")

	// ErrWhileMakingBucket ...
	ErrWhileMakingBucket = errors.New("error while making bucket")

	// ErrPutObject ...
	ErrPutObject = errors.New("error while putting object to minio")

	// ErrGetObject ...
	ErrGetObject = errors.New("error while getting object from minio")

	// ErrPresignedGetObject ...
	ErrPresignedGetObject = errors.New("error while getting URL of an object")

	// ErrInternal ...
	ErrInternal = errors.New("internal server error")

	// ErrWhileWritingToFile ...
	ErrWhileWritingToFile = errors.New("Error while writing to file")
)
