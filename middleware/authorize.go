package middleware

import "github.com/gin-gonic/gin"

type Authorize struct {}

func (*Authorize) Auth(request *gin.Context)  {
	request.Next()
}