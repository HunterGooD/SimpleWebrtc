package app

import (
	"github.com/HunterGooD/SimpleWebrtc/internal/conf"
	"github.com/HunterGooD/SimpleWebrtc/internal/media"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config *conf.Config
	Rooms  *media.Rooms
}

func NewApp() *App {
	return &App{
		Config: &conf.Config{},
		Rooms:  &media.Rooms{},
	}
}

func (a *App) Start() error {

	if err := a.initConfig(); err != nil {
		return err
	}

	r := a.initRouter()

	return r.Run(a.Config.Addr + ":" + a.Config.Port)
}

func (a *App) initConfig() error {
	return a.Config.Init()
}

func (a *App) initRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/room/create", a.createRoom)
	router.POST("/api/room/:hash/:uuid/delete", a.deleteRoom)
	router.Any("/api/room/:hash/:uid/ws", a.wsHandler)

	return router
}
