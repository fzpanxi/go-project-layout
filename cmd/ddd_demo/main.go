package main

import (
	config "ddd_demo/internal/conf"
	"net/http"
)

type App struct {
	hs *http.Server
}

func newApp(hs *http.Server) *App {
	return &App{
		hs: hs,
	}
}
func (app *App) Start() {
	app.hs.ListenAndServe()
}
func main() {
	bc, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Start()
	//log.Println(userUseCase.GetUser(context.Background(), 0))
}
