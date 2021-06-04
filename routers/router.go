package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/middleware/jwt"
	"github.com/qiaocco/go-gin-example/pkg/settings"
	v1 "github.com/qiaocco/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(settings.RunMode)

	r.GET("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")
	// 引入jwt中间件
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags/:id", v1.EditTags)
		apiv1.DELETE("/tags/:id", v1.DeleteTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
