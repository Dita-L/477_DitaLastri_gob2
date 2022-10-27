package router

import (
	"final/controllers"
	"final/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.POST("/", middlewares.Authentication(), controllers.CreatePhoto)
		photoRouter.GET("/", middlewares.Authentication(), controllers.GetAllPhoto)
		photoRouter.PUT("/:photoid", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.Updatephoto)
		photoRouter.DELETE("/:photoid", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.DeletePhoto)

	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", middlewares.Authentication(), controllers.CreateComment)
		commentRouter.PUT("/:commentid", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentid", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socmedRouter := r.Group("/socialmedias")
	{
		socmedRouter.POST("/", middlewares.Authentication(), controllers.CreateSocMed)
		socmedRouter.PUT("/:socialmediaid", middlewares.Authentication(), middlewares.SocmedAuthorization(), controllers.Updatesocmed)
		socmedRouter.DELETE("/:socialmediaid", middlewares.Authentication(), middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}
	return r
}
