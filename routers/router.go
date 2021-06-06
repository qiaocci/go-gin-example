package routers

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/qiaocco/go-gin-example/docs"
	"github.com/qiaocco/go-gin-example/middleware/jwt"
	"github.com/qiaocco/go-gin-example/pkg/setting"
	v1 "github.com/qiaocco/go-gin-example/routers/api/v1"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: setting.AppSetting.SentryDsn,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	setSwaggerInfo()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// Once it's done, you can attach the handler as one of your middleware
	r.Use(sentrygin.New(sentrygin.Options{}))

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func setSwaggerInfo() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "go gin example api"
	docs.SwaggerInfo.Description = "This is a blog server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "qiaocco.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
