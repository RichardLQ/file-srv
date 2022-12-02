package main

import (
	"github.com/RichardLQ/confs"
	"github.com/RichardLQ/file-srv/route"
	"github.com/RichardLQ/fix-srv/client"
	"github.com/gin-gonic/gin"
)

func main() {
	confs.NewStart().BinComb(&client.Global)
	e := gin.Default()
	route.IndexRouter(e)
	e.Run(client.Global.FileServiceConf.HttpIp+":"+
		client.Global.FileServiceConf.HttpPort)
}

