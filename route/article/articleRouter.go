package article

import (
	"github.com/gin-gonic/gin"
	"go-blog/server/controller/article"
)

type ArticleRouter struct {}

func (*ArticleRouter) Init(r gin.IRoutes)  {
	ctl := new(article.ArticleController)
	r.GET("article", ctl.List)
	r.GET("article:id", ctl.Info)
}