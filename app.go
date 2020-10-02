package main

import (
	"flag"
	"strings"
	"webssh/controller/ws"
	"webssh/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := router.InitRouter()
	go ws.Manager.Start()
	router.Run(strings.Join([]string{":", port}, ""))
}
