package ports

import "gorm.io/gorm"

type IDBManager interface {
	GetConnection() (*gorm.DB, error)
	EnsureConnection() error
	Close() error
}
