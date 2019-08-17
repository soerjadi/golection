package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/soerjadi/golection/utils"
)

func main() {

	var hostIP = utils.GetEnv("HOST_IP", "127.0.0.1")
	var port = utils.GetEnv("PORT", "8080")
	var debug = utils.GetEnv("DEBUG", "true")

	var app = gin.Default()

	if debug == "false" {
		gin.SetMode(gin.ReleaseMode)
	}

	initializeRoutes(app)

	_ = app.Run(fmt.Sprintf("%s:%s", hostIP, port))

}

func initializeRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":  200,
			"data":  "hello world",
			"error": "",
		})
	})
}
