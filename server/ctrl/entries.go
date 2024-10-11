package ctrl

import (
	"errors"
	"gopherss/db"
	"gopherss/httputil"
	"gopherss/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntriesController struct{}

var entriesRepository = db.EntriesRepository{}

func (c EntriesController) Register(r *gin.RouterGroup) {
	group := r.Group("/entry")

	group.GET("/", c.list)
	group.PUT("/", c.updateMany)
	group.POST("/mark", c.markMany)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.POST("/:id/mark", c.mark)
}

// list godoc
//
//	@Summary		List entries
//	@Description	get entries
//	@Tags			entry
//	@Produce		json
//	@Param			query		query	model.ListEntriesInput	false	"Query"
//	@Success		200			{array}	model.Entry
//	@Router			/entry		[get]
func (c EntriesController) list(ctx *gin.Context) {
	query := model.ListEntriesInput{}
	err := ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	entries, err := entriesRepository.GetMany(query)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, entries)
}

// updateMany godoc
//
//	@Summary		Update entries
//	@Description	update entries
//	@Tags			entry
//	@Accept			json
//	@Produce		json
//	@Param			entries	body	model.UpdateEntriesInput	true	"Update entries"
//	@Success		204
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/entry	[put]
func (c EntriesController) updateMany(ctx *gin.Context) {
	body := model.UpdateEntriesInput{}
	err := ctx.BindJSON(&body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = entriesRepository.UpdateMany(body)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// markMany godoc
//
//	@Summary		Mark as read/unread
//	@Description	mark entries as read/unread up to a timestamp
//	@Tags			entry
//	@Param			query		query	model.MarkEntriesInput	false	"Query"
//	@Success		204
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/mark	[post]
func (c EntriesController) markMany(ctx *gin.Context) {
	query := model.MarkEntriesInput{}
	err := ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = entriesRepository.MarkMany(query)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// get godoc
//
//	@Summary		Get entry
//	@Description	get entry
//	@Tags			entry
//	@Produce		json
//	@Param        	id   		path      	int	true	"Entry ID"
//	@Success		200			{object}	model.Entry
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/{id}	[get]
func (c EntriesController) get(ctx *gin.Context) {
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
	entry, err := entriesRepository.Get(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	if entry == nil {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("not found"))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

// update godoc
//
//	@Summary		Update entry
//	@Description	update entry
//	@Tags			entry
//	@Accept			json
//	@Produce		json
//	@Param        	id   		path      	int						true	"Entry ID"
//	@Param			entry		body		model.UpdateEntryInput	true	"Update entry"
//	@Success		200			{object}	model.Entry
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/{id}	[put]
func (c EntriesController) update(ctx *gin.Context) {
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
	body := model.UpdateEntryInput{}
	err = ctx.BindJSON(&body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	updated, err := entriesRepository.Update(id, body)
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

// mark godoc
//
//	@Summary		Mark as read/unread
//	@Description	mark entry as read/unread
//	@Tags			entry
//	@Param        	id  				path    	int						true	"Entry ID"
//	@Param			query				query		model.MarkEntryInput	false	"Query"
//	@Success		204
//	@Failure		400					{object}	httputil.HTTPError
//	@Failure		404					{object}	httputil.HTTPError
//	@Router			/entry/{id}/mark	[post]
func (c EntriesController) mark(ctx *gin.Context) {
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
	query := model.MarkEntryInput{}
	err = ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	updated, err := entriesRepository.Mark(id, query)
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
