package main

import (
	"giligili/config"
	"giligili/server"
)

func main()  {

	config.Init()

	server.Router()
}
