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
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	supadminHandler *handler.SupAdminHandler) *ServerHTTP {

	engine := gin.Default()

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.PATCH("/forgotpass", userHandler.ForgotPassword)

		model := user.Group("/model")
		{
			model.GET("/", productHandler.ListAllModel)
			model.GET("/:id", productHandler.ListModel)
		}
		brands := user.Group("/brands")
		{
			brands.GET("/", productHandler.ListAllBrand)
			brands.GET("/:id", productHandler.ListBrand)
		}
		category := user.Group("/category")
		{
			category.GET("/", productHandler.ListAllCategories)
			category.GET("/:id", productHandler.ListCategory)
		}

		user.Use(middleware.UserAuth)
		{
			user.POST("/logout", userHandler.UserLogout)

			profile := user.Group("/profile")
			{
				profile.GET("/", userHandler.ViewProfile)
				profile.PATCH("/edit", userHandler.EditProfile)
				profile.PATCH("/updatepassword", userHandler.UpdatePassword)
			}
			address := user.Group("/address")
			{
				address.POST("/add", userHandler.AddAddress)
				address.PATCH("/update/:addressId", userHandler.UpdateAddress)
			}
			cart := user.Group("/cart")
			{
				cart.POST("/add/:model_id", cartHandler.AddToCart)
				cart.PATCH("/remove/:model_id", cartHandler.RemoveFromCart)
				cart.GET("/", cartHandler.ListCart)
			}
			order := user.Group("/order")
			{
				order.POST("/orderall/:payment_id", orderHandler.OrderAll)
				order.PATCH("/cancel/:orderId", orderHandler.UserCancelOrder)
				order.GET("/:orderId", orderHandler.ListOrder)
				order.GET("/", orderHandler.ListAllOrders)

			}
		}

	}
	admin := engine.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)
		admin.Use(middleware.AdminAuth)
		{
			admin.POST("/logout", adminHandler.AdminLogout)

			adminUsers := admin.Group("/user")
			{
				adminUsers.GET("/:user_id", adminHandler.ShowUser)
				adminUsers.GET("/", adminHandler.ShowAllUsers)
			}

			category := admin.Group("/category")
			{
				category.POST("/create", productHandler.CreateCategory)
				category.PATCH("/update/:id", productHandler.UpdateCategory)
				category.DELETE("/delete/:id", productHandler.DeleteCategory)
				category.GET("/", productHandler.ListAllCategories)
				category.GET("/:id", productHandler.ListCategory)
			}

			brand := admin.Group("/brand")
			{
				brand.POST("/create", productHandler.AddBrand)
				brand.PATCH("/update/:id", productHandler.UpdateBrand)
				brand.DELETE("/delete/:id", productHandler.DeleteBrand)
				brand.GET("/", productHandler.ListAllBrand)
				brand.GET("/:id", productHandler.ListBrand)
			}
			model := admin.Group("/model")
			{
				model.POST("/add", productHandler.AddModel)
				model.PATCH("/update/:id", productHandler.UpdateModel)
				model.DELETE("/delete/:id", productHandler.DeleteModel)
				model.GET("/", productHandler.ListAllModel)
				model.GET("/:id", productHandler.ListModel)
				model.POST("/uploadimage/:id", productHandler.UploadImage)

			}
			order := admin.Group("/order")
			{
				order.PATCH("/update", orderHandler.UpdateOrder)
			}
		}

	}
	supadmin := engine.Group("/supadmin")
	{
		supadmin.POST("/login", supadminHandler.SupAdminLogin)
		supadmin.Use(middleware.SupAdminAuth)
		{
			supadmin.POST("/logout", supadminHandler.SupAdminLogout)

			supAdminUsers := supadmin.Group("/user")
			{
				supAdminUsers.PATCH("/block", supadminHandler.BlockUser)
				supAdminUsers.PATCH("/unblock/:user_id", supadminHandler.UnblockUser)
			}
		}

	}

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start() {
	sh.engine.Run()
}
