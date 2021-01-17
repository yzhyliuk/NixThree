package application

import "sync"

//App : exported interface for starting an application
type App interface {
	Start() error
}
//parseApp : internal type that implements App interface
type parseApp struct {
	sourceUrl string
	//global wait group for all goroutines in the app
	wg sync.WaitGroup
}
//InitApp : initialize an new Application returning an App interface instance
func InitApp(sourceUrl string) App {
	app := new(parseApp)
	app.sourceUrl = sourceUrl
	return app
}
//Start : starts new Application that parses comments and posts
func (pa *parseApp) Start() error {
	pa.getPostByUserID(7)
	//waiting for all goroutines to finish
	pa.wg.Wait()
	return nil
}
