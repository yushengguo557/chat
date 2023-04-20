package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yushengguo557/chat/api/v1"
	"github.com/yushengguo557/chat/global"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())   // 使用Logger中间件
	r.Use(gin.Recovery()) // 使用Recovery中间件
	gin.DefaultWriter = global.Logger.Writer()
	// r.Use(gin.LoggerWithWriter(global.Logger.Writer()))
	apiv1 := r.Group("/v1") // 路由组
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/admin", v1.Admin)
		apiv1.POST("/logout", v1.Logout)
		apiv1.PUT("/me/:id/", v1.UpdateMyInfo)
		apiv1.GET("/me/:id/", v1.GetMyInfo)

		apiv1.GET("/me/:id/friends", v1.GetMyFriends)
		apiv1.POST("/friend/:id", v1.AddFriend)
		apiv1.DELETE("/friend/:id", v1.DeleteFriend)
		apiv1.PUT("/friend/:id", v1.UpdateFriendInfo)
		apiv1.GET("/friend/:id", v1.GetFriendInfo)

		apiv1.POST("/message", v1.SendMessage)
		apiv1.DELETE("/message/:id", v1.DeleteMessage)
		apiv1.PUT("/message/:id", v1.UpdateMessage)
		apiv1.GET("/message", v1.ReceiveMessage)
	}
	return r
}
