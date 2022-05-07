package route

import (
	"github.com/gin-gonic/gin"
	"go-blog/route/article"
)

type RouterGroup struct {
	Article article.ArticleRouter
}


func (group *RouterGroup) InitRouter(routers gin.IRoutes) {
	group.Article.Init(routers)
}
