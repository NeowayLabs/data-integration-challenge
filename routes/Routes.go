package routes

import (
	"github.com/coscms/xorm"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	handlers "github.com/ruiblaese/data-integration-challenge/handlers"
	middleware "github.com/ruiblaese/data-integration-challenge/middleware"
)

//StartRouter inicia servidor e estabelece rotas
func StartRouter(router *gin.Engine, xormEngine *xorm.Engine) *gin.Engine {

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.XormMiddleware(xormEngine))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group("/api/v1")

	//rotas para company
	companies := v1.Group("/companies")
	{
		companies.GET("", handlers.GetCompany)
		companies.GET(":id", handlers.GetCompanyByID)
		//companies.GET("id/:id", handlers.GetCompanyByID)
		//companies.GET("name/:name", handlers.GetCompanyByName)
		companies.PUT(":id", handlers.PutCompany)
		companies.POST("", handlers.NewCompany)
		companies.DELETE(":id", handlers.DeleteCompany)
		companies.POST("upload", handlers.UploadCompanysWithCSV)
	}

	return router
}
