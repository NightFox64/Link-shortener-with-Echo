package shortening

import (
	"gorm.io/gorm"
)

var GlobalDB, _ = Setup()

type BaseRepo struct {
	GlobalDB *gorm.DB
}
