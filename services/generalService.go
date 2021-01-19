package services

import (
	"gorm.io/gorm"
)

//Postservice :
type generalservive struct {
	db *gorm.DB
}

//NewGenaralService : Returns new general Service
func NewGenaralService() Service {
	return &generalservive{}
}

//AddDataSource : adds data source for given service
func (s *generalservive) AddDataSource(db *gorm.DB) {
	s.db = db
}

//Create : creates a new entry of row in db
func (s *generalservive) Create(obj interface{}) error {
	result := s.db.Create(obj)
	return result.Error
}

//Recive : gets data from db by primary key
func (s *generalservive) Recive(obj interface{}, primaryKey int) (interface{}, error) {
	result := s.db.First(obj, primaryKey)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

//Update : updates post merging empty fields with current values
func (s *generalservive) Update(obj interface{}) error {
	result := s.db.Save(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete : deletes first entry of object with given id
func (s *generalservive) Delete(obj interface{}, primaryKey int) error {
	result := s.db.Delete(obj, primaryKey)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
