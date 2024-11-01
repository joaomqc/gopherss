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

type FeedsController struct{}

var feedsService = svc.FeedsService{}

func (c FeedsController) Register(r *gin.RouterGroup) {
	group := r.Group("/feed")

	group.GET("/", c.list)
	group.POST("/", c.create)
	group.POST("/refresh", c.refreshMany)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.GET("/:id/icon", c.getIcon)
	group.POST("/:id/refresh", c.refresh)
}

// list godoc
//
//	@Summary		List feeds
//	@Description	get feeds
//	@Tags			feed
//	@Produce		json
//	@Param			query		query	model.ListFeedsInput	false	"Query"
//	@Success		200			{array}	model.Feed
//	@Router			/feed		[get]
func (c FeedsController) list(ctx *gin.Context) {
	query := model.ListFeedsInput{}
	err := ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	feeds, err := feedsService.GetMany(query)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, feeds)
}

// create godoc
//
//	@Summary		Create feed
//	@Description	create feed
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			feed	body		model.AddFeedInput	true	"Add feed"
//	@Success		200		{object}	model.Feed
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/feed	[post]
func (c FeedsController) create(ctx *gin.Context) {
	body := model.AddFeedInput{}
	err := ctx.BindJSON(&body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	id, err := feedsService.Create(body)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// refreshMany godoc
//
//	@Summary		Refresh feeds
//	@Description	refresh many feeds
//	@Tags			feed
//	@Param			category	query	int	false	"Category id"
//	@Success		204
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/feed/refresh	[post]
func (c FeedsController) refreshMany(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// get godoc
//
//	@Summary		Get feed
//	@Description	get feed
//	@Tags			feed
//	@Produce		json
//	@Param        	id   		path      	int	true	"Feed ID"
//	@Success		200			{object}	model.Feed
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/feed/{id}	[get]
func (c FeedsController) get(ctx *gin.Context) {
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
	feed, err := feedsService.Get(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	if feed == nil {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("not found"))
		return
	}
	ctx.JSON(http.StatusOK, feed)
}

// update godoc
//
//	@Summary		Update feed
//	@Description	update feed
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param        	id   		path      	int						true	"Feed ID"
//	@Param			feed		body		model.UpdateFeedInput	true	"Update feed"
//	@Success		200			{object}	model.Feed
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/feed/{id}	[put]
func (c FeedsController) update(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// delete godoc
//
//	@Summary		Delete feed
//	@Description	delete feed
//	@Tags			feed
//	@Produce		json
//	@Param        	id   		path      	int	true	"Feed ID"
//	@Success		204
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/feed/{id}	[delete]
func (c FeedsController) delete(ctx *gin.Context) {
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
	err = feedsService.Delete(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// getIcon godoc
//
//	@Summary		Get feed icon
//	@Description	get feed icon
//	@Tags			feed
//	@Produce		json
//	@Param        	id   			path      	int	true	"Feed ID"
//	@Success		200				{object}	model.Feed
//	@Failure		404				{object}	httputil.HTTPError
//	@Router			/feed/{id}/icon	[get]
func (c FeedsController) getIcon(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// refresh godoc
//
//	@Summary		Refresh feed
//	@Description	refresh feed
//	@Tags			feed
//	@Param        	id   				path      	int	true	"Feed ID"
//	@Success		204
//	@Failure		404					{object}	httputil.HTTPError
//	@Router			/feed/{id}/refresh	[post]
func (c FeedsController) refresh(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}
