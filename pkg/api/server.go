package http

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/api/handler"
	"main.go/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler) *ServerHTTP {

	engine := gin.Default()

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)

		user.Use(middleware.UserAuth)
		{
			user.POST("/logout", userHandler.UserLogout)
		}

	}
	admin := engine.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)
		admin.Use(middleware.AdminAuth)
		{
			admin.POST("/logout", adminHandler.AdminLogout)

			category := admin.Group("/category")
			{
				category.POST("/create", productHandler.CreateCategory)
				category.PATCH("/update/:id", productHandler.UpdateCategory)
				category.DELETE("/delete/:id", productHandler.DeleteCategory)
				category.GET("/listall", productHandler.ListAllCategories)
				category.GET("/list/:id", productHandler.ListCategory)
			}

			brand := admin.Group("/brand")
			{
				brand.POST("/create", productHandler.AddBrand)
				brand.PATCH("/update/:id", productHandler.UpdateBrand)
				brand.DELETE("/delete/:id", productHandler.DeleteBrand)
				brand.GET("/listall", productHandler.ListAllBrand)
				brand.GET("/list/:id", productHandler.ListBrand)
			}
			model := admin.Group("/model")
			{
				model.POST("/add", productHandler.AddModel)

			}
		}

	}

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start() {
	sh.engine.Run()
}
