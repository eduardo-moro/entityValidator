package router

import (
	request "entityValidator.com/request"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("/go/src/entityValidator.com/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.GET("/api/cfn/codigo/:registro", func(c *gin.Context) {
		codigo := c.Param("registro")

		status, body := request.GetCfnByCode(codigo)

		log.Println(status)

		c.JSON(http.StatusOK, gin.H{
			"items": body,
			"error": "",
		})
	})

	return router
}
