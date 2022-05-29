package uuidgenerator

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/sardortoshkentov/minio-example/pkg/errs"
)

func GenerateUUID() (uuid.UUID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error while generating an uuid: ", err.Error())
		return uuid, errs.ErrInternal
	}

	f, err := os.Create("./pkg/objectName.txt")

	if err != nil {
		log.Println("Error while creating a file, error: ", err.Error())
		return uuid, err
	}

	defer f.Close()

	_, err = f.WriteString(uuid.String())
	if err != nil {
		log.Println("Error while writing to file: ", err.Error())
		return uuid, errs.ErrWhileWritingToFile
	}

	return uuid, nil
}
