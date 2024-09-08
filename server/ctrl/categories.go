package ctrl

import "github.com/gin-gonic/gin"

type CategoriesController struct{}

func (c CategoriesController) Register(r *gin.Engine) {
	group := r.Group("/category")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.PUT("/", c.update)
	group.DELETE("/", c.delete)
}

func (c CategoriesController) getAll(ctx *gin.Context) {

}

func (c CategoriesController) create(ctx *gin.Context) {

}

func (c CategoriesController) update(ctx *gin.Context) {

}

func (c CategoriesController) delete(ctx *gin.Context) {

}
