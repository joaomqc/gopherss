package ctrl

import "github.com/gin-gonic/gin"

type CategoriesController struct{}

func (c CategoriesController) Register(r *gin.Engine) {
	group := r.Group("/category")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.POST("/:id/read", c.markAsRead)
	group.POST("/:id/refresh", c.refresh)
}

func (c CategoriesController) getAll(ctx *gin.Context) {

}

func (c CategoriesController) create(ctx *gin.Context) {

}

func (c CategoriesController) update(ctx *gin.Context) {

}

func (c CategoriesController) delete(ctx *gin.Context) {

}

func (c CategoriesController) markAsRead(ctx *gin.Context) {
	// ?before={timestamp} mark as read/unread up to timestamp
}

func (c CategoriesController) refresh(ctx *gin.Context) {
	// refresh all feeds in category
}
