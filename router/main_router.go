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

	router.GET("/api/main", func(c *gin.Context) {
		status, body := request.GetCfn()

		log.Println(status)

		c.JSON(http.StatusOK, gin.H{
			"body": body,
		})
	})

	return router
}
