package ctrl

import (
	"errors"
	"gopherss/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct{}

func (c CategoriesController) Register(r *gin.RouterGroup) {
	group := r.Group("/category")

	group.GET("/", c.list)
	group.POST("/", c.create)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.POST("/:id/mark", c.mark)
	group.POST("/:id/refresh", c.refresh)
}

// @Summary		List categories
// @Description	get categories
// @Tags			category
// @Produce		json
// @Success		200			{array}	model.Category
// @Router			/category	[get]
func (c CategoriesController) list(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// @Summary		Create category
// @Description	create category
// @Tags			category
// @Accept			json
// @Produce		json
// @Param			category	body		model.AddCategory	true	"Add category"
// @Success		200			{object}	model.Category
// @Failure		400			{object}	httputil.HTTPError
// @Router			/category	[post]
func (c CategoriesController) create(ctx *gin.Context) {

}

// @Summary		Get category
// @Description	get category
// @Tags			category
// @Produce		json
// @Success		200				{object}	model.Category
// @Failure		404				{object}	httputil.HTTPError
// @Router			/category/{id}	[get]
func (c CategoriesController) get(ctx *gin.Context) {

}

// @Summary		Update category
// @Description	update category
// @Tags			category
// @Accept			json
// @Produce		json
// @Param			category		body		model.UpdateCategory	true	"Update category"
// @Success		200				{object}	model.Category
// @Failure		400				{object}	httputil.HTTPError
// @Failure		404				{object}	httputil.HTTPError
// @Router			/category/{id}	[put]
func (c CategoriesController) update(ctx *gin.Context) {

}

// @Summary		Delete category
// @Description	delete category
// @Tags			category
// @Produce		json
// @Success		204
// @Failure		404				{object}	httputil.HTTPError
// @Router			/category/{id}	[delete]
func (c CategoriesController) delete(ctx *gin.Context) {

}

// @Summary		Mark as read/unread
// @Description	mark category entries as read/unread up to a timestamp
// @Tags			category
// @Param			before	query	string			true	"Timestamp to mark/unread as read to"	format(date-time)	example("2006-01-02 15:04:05")
// @Param			as		query	model.Status	true	"New status"
// @Success		204
// @Failure		400					{object}	httputil.HTTPError
// @Failure		404					{object}	httputil.HTTPError
// @Router			/category/{id}/mark	[post]
func (c CategoriesController) mark(ctx *gin.Context) {

}

// @Summary		Refresh feeds
// @Description	refresh all feeds in the category
// @Tags			category
// @Success		204
// @Failure		404						{object}	httputil.HTTPError
// @Router			/category/{id}/refresh	[post]
func (c CategoriesController) refresh(ctx *gin.Context) {
	// refresh all feeds in category
}
