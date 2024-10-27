package ctrl

import (
	"errors"
	"gopherss/httputil"
	"gopherss/model"
	"gopherss/svc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct{}

var categoriesService = svc.CategoriesService{}

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
//	@Param			query		query	model.ListCategoriesInput	false	"Query"
//	@Success		200			{array}	model.Category
//	@Router			/category	[get]
func (c CategoriesController) list(ctx *gin.Context) {
	query := model.ListCategoriesInput{}
	err := ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	categories, err := categoriesService.GetMany(query)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// create godoc
//
//	@Summary		Create category
//	@Description	create category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		model.AddCategoryInput	true	"Add category"
//	@Success		200			{object}	model.Category
//	@Failure		400			{object}	httputil.HTTPError
//	@Router			/category	[post]
func (c CategoriesController) create(ctx *gin.Context) {
	body := model.AddCategoryInput{}
	err := ctx.BindJSON(&body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = categoriesService.Create(body)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusCreated)
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
	idStr := ctx.Param("id")
	if idStr == "" {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is mandatory"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is invalid"))
		return
	}
	category, err := categoriesService.Get(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	if category == nil {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("not found"))
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// update godoc
//
//	@Summary		Update category
//	@Description	update category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param        	id   			path      	int							true	"Category ID"
//	@Param			category		body		model.UpdateCategoryInput	true	"Update category"
//	@Success		200				{object}	model.Category
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/category/{id}	[put]
func (c CategoriesController) update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is mandatory"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is invalid"))
		return
	}
	body := model.UpdateCategoryInput{}
	err = ctx.BindJSON(&body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	updated, err := categoriesService.Update(id, body)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	if !updated {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("not found"))
		return
	}
	ctx.Status(http.StatusNoContent)
}

// delete godoc
//
//	@Summary		Delete category
//	@Description	delete category
//	@Tags			category
//	@Produce		json
//	@Param        	id   			path      	int		true	"Category ID"
//	@Param        	keepFeeds		path      	bool	false	"Keep feeds"
//	@Success		204
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/category/{id}	[delete]
func (c CategoriesController) delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is mandatory"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("id is invalid"))
		return
	}
	keepFeedsStr := ctx.Param("keepFeeds")
	keepFeeds, err := strconv.ParseBool(keepFeedsStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("keepFeeds is invalid"))
		return
	}
	err = categoriesService.Delete(id, keepFeeds)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
