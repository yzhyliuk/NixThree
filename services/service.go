package services

import "gorm.io/gorm"

//Service : defines an interface of Services
type Service interface {
	AddDataSource(db *gorm.DB)
	Create(obj interface{}) error
	Recive(obj interface{}, primaryKey int) (interface{}, error)
	Update(obj interface{}) error
	Delete(obj interface{}, primaryKey int) error
}
