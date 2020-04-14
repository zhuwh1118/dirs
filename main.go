package main

import (
	"flag"
	"zwh/web"

	"github.com/gin-gonic/gin"
)

func init() {
	flag.StringVar(&web.Root, "root", "/", "the root path")
}

func main() {
	flag.Parse()
	r := gin.Default()
	r.GET("/v1/dirs", web.Dirs)
	r.Run(":9000")
}
