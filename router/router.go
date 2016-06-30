package router

import (
	"github.com/shimastripe/gouserapi/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("api")
	{
		//USER API
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

		//PROFILE API
		api.GET("/profiles", controllers.GetProfiles)
		api.GET("/profiles/:id", controllers.GetProfile)
		api.POST("/profiles", controllers.CreateProfile)
		api.PUT("/profiles/:id", controllers.UpdateProfile)
		api.DELETE("/profiles/:id", controllers.DeleteProfile)

		//ACCOUNT_NAME API
		api.GET("/account_names", controllers.GetAccountNames)
		api.GET("/account_names/:id", controllers.GetAccountName)
		api.POST("/account_names", controllers.CreateAccountName)
		api.PUT("/account_names/:id", controllers.UpdateAccountName)
		api.DELETE("/account_names/:id", controllers.DeleteAccountName)

		//EMAIL API
		api.GET("/emails", controllers.GetEmails)
		api.GET("/emails/:id", controllers.GetEmail)
		api.POST("/emails", controllers.CreateEmail)
		api.PUT("/emails/:id", controllers.UpdateEmail)
		api.DELETE("/emails/:id", controllers.DeleteEmail)
	}
}
