package routers

import (
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	//注册中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()

	upload := api.NewUpload()

	r.POST("/upload/file", upload.UploadFile)

	//访问静态资源
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", v1.GetAuth)

	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT()) //某个路由分组需要使用jwt

	{
		//tags
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)
		apiv1.GET("/tags/abc", tag.Get)

		//article
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.List)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}

/*
curl -X POST http://127.0.0.1:9999/upload/file -F file=@hello.jpeg -F  type=2
{"code":2003001,"details":["file suffix is not supported"],"msg":"上传文件失败"}
*/
