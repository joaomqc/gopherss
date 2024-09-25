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
}

// list godoc
//
//	@Summary		List categories
//	@Description	get categories
//	@Tags			category
//	@Produce		json
//	@Success		200			{array}	model.Category
//	@Router			/category	[get]
func (c CategoriesController) list(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// create godoc
//
//	@Summary		Create category
//	@Description	create category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		model.AddCategory	true	"Add category"
//	@Success		200			{object}	model.Category
//	@Failure		400			{object}	httputil.HTTPError
//	@Router			/category	[post]
func (c CategoriesController) create(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// get godoc
//
//	@Summary		Get category
//	@Description	get category
//	@Tags			category
//	@Produce		json
//	@Param        	id   			path      	int	true	"Category ID"
//	@Success		200				{object}	model.Category
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/category/{id}	[get]
func (c CategoriesController) get(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// update godoc
//
//	@Summary		Update category
//	@Description	update category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param        	id   			path      	int						true	"Category ID"
//	@Param			category		body		model.UpdateCategory	true	"Update category"
//	@Success		200				{object}	model.Category
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/category/{id}	[put]
func (c CategoriesController) update(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// delete godoc
//
//	@Summary		Delete category
//	@Description	delete category
//	@Tags			category
//	@Produce		json
//	@Param        	id   			path      	int	true	"Category ID"
//	@Success		204
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/category/{id}	[delete]
func (c CategoriesController) delete(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}
