package router

import (
	request "entityValidator.com/request"
	gin "github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("/go/src/entityValidator.com/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	api := router.Group("/api")
		api.GET("/cnpj/:cnpj", func(c *gin.Context) {
			cnpj := c.Param("cnpj")
			body := request.GetCnpj(cnpj)
			c.Data(http.StatusOK, "application/json", []byte(body))
		})

		api.GET("/cref/:cref", func(c *gin.Context) {
			cref := c.Param("cref")
			body := request.GetCref(cref)
			c.Data(http.StatusOK, "application/json", []byte(body))
		})

		api.GET("/crefpj/:cref", func(c *gin.Context) {
			cref := c.Param("cref")
			body := request.GetCrefPj(cref)
			c.Data(http.StatusOK, "application/json", []byte(body))
		})

	cfn := api.Group("/cfn")
		cfn.GET("/codigo/:registro", func(c *gin.Context) {
			codigo := c.Param("registro")
			c.Data(http.StatusOK, "application/json", request.GetCfnByCode(codigo))
		})

		cfn.GET("/nome/:nome", func(c *gin.Context) {
			nome := c.Param("nome")
			body := request.GetCfnByName(nome)
			c.Data(http.StatusOK, "application/json", []byte(body))
		})


	return router
}
