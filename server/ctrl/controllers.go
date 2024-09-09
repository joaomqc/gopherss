package ctrl

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(*gin.RouterGroup)
}

type BaseQuery struct {
	Start int    `form:"start"`
	End   int    `form:"end"`
	Order string `form:"order"`
	Sort  string `form:"sort"`
}
