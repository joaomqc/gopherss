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
	group.POST("/", c.create)
	group.PUT("/", c.updateMany)
	group.POST("/mark", c.markALl)
	group.GET("/:id", c.get)
	group.PUT("/:id", c.update)
	group.DELETE("/:id", c.delete)
	group.POST("/:id/star", c.star)
	group.POST("/:id/mark", c.mark)
}

func (c EntriesController) list(ctx *gin.Context) {
	// &(feed|category)=&limit=&starred=&read=&search=&offset=
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

func (c EntriesController) create(ctx *gin.Context) {

}

func (c EntriesController) updateMany(ctx *gin.Context) {
	// body: {"entries": [1,2,3],"status":"read"}
}

func (c EntriesController) markALl(ctx *gin.Context) {
	// ?before=&as= mark as read/unread up to timestamp
}

func (c EntriesController) get(ctx *gin.Context) {

}

func (c EntriesController) update(ctx *gin.Context) {

}

func (c EntriesController) delete(ctx *gin.Context) {

}

func (c EntriesController) star(ctx *gin.Context) {

}

func (c EntriesController) mark(ctx *gin.Context) {
	// ?as=
}
