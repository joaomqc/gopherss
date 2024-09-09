package ctrl

import (
	"gopherss/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EntriesController struct{}

var entriesRepository = db.EntriesRepository{}

type Query struct {
	BaseQuery
	Category string `form:"category"`
	Feed     string `form:"feed"`
}

func (c EntriesController) Register(r *gin.RouterGroup) {
	group := r.Group("/entry")

	group.GET("/", c.list)
	group.PUT("/", c.updateMany)
	group.POST("/mark", c.markAll)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.POST("/:id/mark", c.mark)
}

// @Summary		List entries
// @Description	get entries
// @Tags			entry
// @Produce		json
// @Param			category	query	int			false	"Category id"
// @Param			feed	query	int			false	"Feed id"
// @Param			starred	query	bool			false	"Show starred only"
// @Param			read	query	bool			false	"Show read/unread only"
// @Param			search	query	string			false	"Search text"
// @Param			limit	query	int			false	"Max entries to return"
// @Param			offset	query	int			false	"Query offset"
// @Success		200			{array}	model.Entry
// @Router			/entry	[get]
func (c EntriesController) list(ctx *gin.Context) {
	query := Query{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	entries, err := entriesRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entries)
}

// @Summary		Update entries
// @Description	update entries
// @Tags			entry
// @Accept			json
// @Produce		json
// @Param			entries	body		model.UpdateEntries	true	"Update entries"
// @Success		204
// @Failure		400			{object}	httputil.HTTPError
// @Router			/entry	[put]
func (c EntriesController) updateMany(ctx *gin.Context) {

}

// @Summary		Mark as read/unread
// @Description	mark entries as read/unread up to a timestamp
// @Tags			entry
// @Param			before	query	string			true	"Timestamp to mark/unread as read to"	format(date-time)	example("2006-01-02 15:04:05")
// @Param			as		query	model.Status	true	"New status"
// @Success		204
// @Failure		400					{object}	httputil.HTTPError
// @Failure		404					{object}	httputil.HTTPError
// @Router			/entry/mark	[post]
func (c EntriesController) markAll(ctx *gin.Context) {

}

// @Summary		Get entry
// @Description	get entry
// @Tags			entry
// @Produce		json
// @Success		200				{object}	model.Entry
// @Failure		404				{object}	httputil.HTTPError
// @Router			/entry/{id}	[get]
func (c EntriesController) get(ctx *gin.Context) {

}

// @Summary		Update entry
// @Description	update entry
// @Tags			entry
// @Accept			json
// @Produce		json
// @Param			entry		body		model.UpdateEntry	true	"Update entry"
// @Success		200				{object}	model.Entry
// @Failure		400				{object}	httputil.HTTPError
// @Failure		404				{object}	httputil.HTTPError
// @Router			/entry/{id}	[put]
func (c EntriesController) update(ctx *gin.Context) {

}

// @Summary		Delete entry
// @Description	delete entry
// @Tags			entry
// @Produce		json
// @Success		204
// @Failure		404				{object}	httputil.HTTPError
// @Router			/entry/{id}	[delete]
func (c EntriesController) delete(ctx *gin.Context) {

}

// @Summary		Mark as read/unread
// @Description	mark entry as read/unread
// @Tags			entry
// @Param			as		query	model.Status	true	"New status"
// @Success		204
// @Failure		400					{object}	httputil.HTTPError
// @Failure		404					{object}	httputil.HTTPError
// @Router			/entry/{id}/mark	[post]
func (c EntriesController) mark(ctx *gin.Context) {

}
