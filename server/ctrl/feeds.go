package ctrl

import "github.com/gin-gonic/gin"

type FeedsController struct{}

func (c FeedsController) Register(r *gin.Engine) {
	group := r.Group("/feed")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.POST("/refresh", c.refreshAll)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.POST("/:id/read", c.markAsRead)
	group.POST("/:id/refresh", c.refresh)
}

func (c FeedsController) getAll(ctx *gin.Context) {
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

func (c FeedsController) markAsRead(ctx *gin.Context) {
	// ?before= mark as read/unread up to timestamp
}

func (c FeedsController) refresh(ctx *gin.Context) {

}
