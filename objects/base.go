package objects

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}