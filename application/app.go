package application

import (
	"NixTwo/services"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorm.io/gorm"
)

//App : exported interface for starting an application
type App interface {
	Start()
}

type webApp struct {
	dataSource     *gorm.DB
	webServer      *http.Server
	webRouter      *http.ServeMux
	generalService services.Service
}

//Start : start function for server
func (w *webApp) Start() {
	go func() {
		w.webServer.ErrorLog.Printf("Starting server on port %s \n", w.webServer.Addr)
		err := w.webServer.ListenAndServe()
		if err != nil {
			w.webServer.ErrorLog.Printf("Error starting server: %s", err.Error())
			os.Exit(1)
		}
	}()
	//Create channel to communicate with os catching terminate commands
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Kill)
	signal.Notify(signChan, os.Interrupt)

	//Block next part of code via wating response from channel
	_ = <-signChan
	//if recieving one - log command and Gracefully shutDown the Server with Timeout of 30 sec
	w.webServer.ErrorLog.Printf("Recived terminate command, graceful shutdown")

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	w.webServer.Shutdown(tc)

}

//NewApp : returns an new App interface with default settings
func NewApp() (App, error) {
	//creating new webApp with default config
	app, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	return app, nil
}
