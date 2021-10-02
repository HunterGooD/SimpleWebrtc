package main

import (
	"github.com/HunterGooD/SimpleWebrtc/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := app.NewApp().Start(); err != nil {
		panic(err)
	}
}
