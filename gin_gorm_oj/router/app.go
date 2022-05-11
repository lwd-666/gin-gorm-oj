package router

import (
	_ "go_code/gin_gorm_oj/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"go_code/gin_gorm_oj/service"
)

func Router() *gin.Engine {
	r:=gin.Default()
	//swager相关配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	//用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login",service.Login)
	r.POST("/send-code",service.SendCode)
	//排行榜
	r.GET("/rank-list", service.GetRankList)
	//提交记录
	r.GET("/submit-list", service.GetSubmitList)
	return r
}
