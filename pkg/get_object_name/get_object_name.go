package getobjectname

import (
	"os"

	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

// GetObjectName ...
func GetObjectName() (string, error) {
	objectName, err := os.ReadFile("./pkg/objectName.txt")
	if err != nil {
		return "", errs.ErrInternal
	}

	return string(objectName), nil
}