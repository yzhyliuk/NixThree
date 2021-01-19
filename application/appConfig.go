package application

import (
	"NixTwo/data"
	"NixTwo/dataSources/mysql"
	"NixTwo/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

var defaultDataBaseName = "blogbase"

var defaultLogger = log.New(os.Stdout, "blog-api: ", log.LstdFlags)

var defaultAppServer = http.Server{
	Addr:     ":9090",
	ErrorLog: defaultLogger,
	Handler:  http.DefaultServeMux,
}

//Returns webApp with default config
func defaultConfig() (*webApp, error) {
	//creating new webApp with default server, router and service
	app := &webApp{
		webServer:      &defaultAppServer,
		webRouter:      http.DefaultServeMux,
		generalService: services.NewGenaralService(),
	}
	//Connecting to db
	db, err := dbDefaultConfig()
	if err != nil {
		return nil, fmt.Errorf("Can't apply default configuration: %s", err.Error())
	}
	//Setting recived gorm.DB as datasource for our app
	app.dataSource = db
	app.generalService.AddDataSource(db)
	app.mapRoutes()
	return app, nil
}

func dbDefaultConfig() (*gorm.DB, error) {
	err := mysql.InitDBConnection(defaultDataBaseName)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database")
	}

	err = mysql.DataBase.AutoMigrate(&data.Post{})
	if err != nil {
		return nil, fmt.Errorf("Can't create a schema for Post struct")
	}
	err = mysql.DataBase.AutoMigrate(&data.Comment{})
	if err != nil {
		return nil, fmt.Errorf("Can't create a schema for Comment struct")
	}
	return mysql.DataBase, nil
}

func (w *webApp) mapRoutes() {
	w.webRouter.HandleFunc("/posts/", w.getPost)
}
