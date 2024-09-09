package ctrl

import "github.com/gin-gonic/gin"

type FeedsController struct{}

func (c FeedsController) Register(r *gin.RouterGroup) {
	group := r.Group("/feed")

	group.GET("/", c.list)
	group.POST("/", c.create)
	group.POST("/refresh", c.refreshAll)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.GET("/:id/icon", c.getIcon)
	group.POST("/:id/mark", c.mark)
	group.POST("/:id/refresh", c.refresh)
}

func (c FeedsController) list(ctx *gin.Context) {
	// ?category=
}

func (c FeedsController) create(ctx *gin.Context) {

}

func (c FeedsController) refreshAll(ctx *gin.Context) {

}

func (c FeedsController) get(ctx *gin.Context) {

}

func (c FeedsController) update(ctx *gin.Context) {

}

func (c FeedsController) delete(ctx *gin.Context) {

}

func (c FeedsController) getIcon(ctx *gin.Context) {

}

func (c FeedsController) mark(ctx *gin.Context) {
	// ?before=&as= mark as read/unread up to timestamp
}

func (c FeedsController) refresh(ctx *gin.Context) {

}
