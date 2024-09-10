package ctrl

import (
	"errors"
	"gopherss/db"
	"gopherss/httputil"
	"gopherss/model"
	"net/http"

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
	group.DELETE("/:id", c.delete)
	group.POST("/:id/mark", c.mark)
}

// list godoc
//
//	@Summary		List entries
//	@Description	get entries
//	@Tags			entry
//	@Produce		json
//	@Param			category	query	int				false	"Category id"
//	@Param			feed		query	int				false	"Feed id"
//	@Param			starred		query	bool			false	"Show starred only"
//	@Param			read		query	bool			false	"Show read/unread only"
//	@Param			search		query	string			false	"Search text"
//	@Param			offset		query	int				false	"Query offset"
//	@Param			limit		query	int				false	"Max entries to return"
//	@Param			order		query	string			false	"Property to order by"
//	@Param			sort		query	model.SortType	false	"Sort ascending/descending"
//	@Success		200			{array}	model.Entry
//	@Router			/entry																																																																		[get]
func (c EntriesController) list(ctx *gin.Context) {
	query := model.EntryListQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	entries, err := entriesRepository.GetAll(query)
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
//	@Param			entries	body	model.UpdateEntries	true	"Update entries"
//	@Success		204
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/entry	[put]
func (c EntriesController) updateMany(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// markMany godoc
//
//	@Summary		Mark as read/unread
//	@Description	mark entries as read/unread up to a timestamp
//	@Tags			entry
//	@Param			category	query	int				false	"Category id"
//	@Param			feed		query	int				false	"Feed id"
//	@Param			before		query	string			true	"Timestamp to mark/unread as read to"	format(date-time)	example("2006-01-02 15:04:05")
//	@Param			as			query	model.Status	true	"New status"
//	@Success		204
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/mark	[post]
func (c EntriesController) markMany(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// get godoc
//
//	@Summary		Get entry
//	@Description	get entry
//	@Tags			entry
//	@Produce		json
//	@Success		200			{object}	model.Entry
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/{id}	[get]
func (c EntriesController) get(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// update godoc
//
//	@Summary		Update entry
//	@Description	update entry
//	@Tags			entry
//	@Accept			json
//	@Produce		json
//	@Param			entry		body		model.UpdateEntry	true	"Update entry"
//	@Success		200			{object}	model.Entry
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/{id}	[put]
func (c EntriesController) update(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// delete godoc
//
//	@Summary		Delete entry
//	@Description	delete entry
//	@Tags			entry
//	@Produce		json
//	@Success		204
//	@Failure		404			{object}	httputil.HTTPError
//	@Router			/entry/{id}	[delete]
func (c EntriesController) delete(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}

// mark godoc
//
//	@Summary		Mark as read/unread
//	@Description	mark entry as read/unread
//	@Tags			entry
//	@Param			as	query	model.Status	true	"New status"
//	@Success		204
//	@Failure		400					{object}	httputil.HTTPError
//	@Failure		404					{object}	httputil.HTTPError
//	@Router			/entry/{id}/mark	[post]
func (c EntriesController) mark(ctx *gin.Context) {
	httputil.NewError(ctx, http.StatusNotImplemented, errors.New("not implemented"))
}
