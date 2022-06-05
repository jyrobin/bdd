package bdd

import "github.com/google/uuid"

const (
	SmallUuidLen  = 13
	MediumUuidLen = 18
)

func Uuid(size ...int) string {
	id := uuid.New().String()
	if len(size) > 0 {
		if sz := size[0]; sz > 0 && sz < len(id) {
			return id[0:sz]
		}
	}
	return id
}

func SmallUuid() string {
	return Uuid(SmallUuidLen)
}

func MediumUuid() string {
	return Uuid(MediumUuidLen)
}
