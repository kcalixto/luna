package serviceUtils

import "github.com/google/uuid"

type ServiceUtils struct {
	GenerateUUID func() string
}

func NewServiceUtils() *ServiceUtils {
	return &ServiceUtils{
		GenerateUUID: generateUUID,
	}
}

func generateUUID() string {
	return uuid.New().String()
}
