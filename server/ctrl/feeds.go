package ctrl

import (
	"errors"
	"gopherss/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedsController struct{}

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
//	@Param			category	query	int	false	"Category id"
//	@Success		200			{array}	model.Feed
//	@Router			/feed																																																																[get]
func (c FeedsController) list(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// create godoc
//
//	@Summary		Create feed
//	@Description	create feed
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			feed	body		model.AddFeed	true	"Add feed"
//	@Success		200		{object}	model.Feed
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/feed	[post]
func (c FeedsController) create(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// refreshMany godoc
//
//	@Summary		Refresh feeds
//	@Description	refresh all feeds
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
//	@Success		200			{object}	model.Feed
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/feed/{id}	[get]
func (c FeedsController) get(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// update godoc
//
//	@Summary		Update feed
//	@Description	update feed
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			feed		body		model.UpdateFeed	true	"Update feed"
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
//	@Success		204
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/feed/{id}	[delete]
func (c FeedsController) delete(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// getIcon godoc
//
//	@Summary		Get feed icon
//	@Description	get feed icon
//	@Tags			feed
//	@Produce		json
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
//	@Success		204
//	@Failure		404					{object}	httputil.HTTPError
//	@Router			/feed/{id}/refresh	[post]
func (c FeedsController) refresh(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}
