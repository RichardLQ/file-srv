package main

import (
	"github.com/RichardLQ/confs"
	"github.com/RichardLQ/file-srv/auth"
	"github.com/RichardLQ/file-srv/route"
	"github.com/gin-gonic/gin"
)

func main() {
	confs.NewStart().BinComb(&auth.Global)
	e := gin.Default()
	route.IndexRouter(e)
	e.Run(auth.Global.FileServiceConf.HttpIp+":"+
		auth.Global.FileServiceConf.HttpPort)
}

