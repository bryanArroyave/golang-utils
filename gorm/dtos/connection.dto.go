package dtos

import "gorm.io/gorm"

type ConnectionDTO struct {
	URI        string
	Env        string
	MaxRetries int
}

type GormConnectionDTO struct {
	*ConnectionDTO
	Dialector *gorm.Dialector
}
