package managers

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/Prabandham/larks_go/db"
	. "github.com/Prabandham/larks_go/objects"
)

func GetObjectByID(id string) (*Project, error) {
	db := db.GetConnection()
	object := Project{}
	uuid, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New("Ivalid ID Format")
	}
	db.Connection.Where("id = ?", uuid).First(&object)
	return &object, nil
}
