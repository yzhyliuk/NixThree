package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     string = "root"
	password string = "s0meG0od@ndStr0ngP@ssWird"
)

var DataBase *gorm.DB

//InitDBConnection : initialize connection to database
func InitDBConnection(dbName string) error {
	db, err := gorm.Open(mysql.Open(getDSN(dbName)), &gorm.Config{})
	if err != nil {
		return err
	}
	DataBase = db
	return nil
}

func getDSN(dbName string) string {
	return fmt.Sprintf("%s:%s@/%s", user, password, dbName)
}
