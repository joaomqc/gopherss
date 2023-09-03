package ctrl

import "github.com/gin-gonic/gin"

type FeedsController struct{}

func (c FeedsController) Register(r *gin.Engine) {
	group := r.Group("/feeds")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.PUT("/", c.update)
	group.DELETE("/", c.delete)
}

func (c FeedsController) getAll(ctx *gin.Context) {

}

func (c FeedsController) create(ctx *gin.Context) {

}

func (c FeedsController) update(ctx *gin.Context) {

}

func (c FeedsController) delete(ctx *gin.Context) {

}
