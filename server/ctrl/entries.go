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

func (c EntriesController) Register(r *gin.Engine) {
	group := r.Group("/entry")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.PUT("/", c.update)
	group.DELETE("/", c.delete)
}

func (c EntriesController) getAll(ctx *gin.Context) {
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

func (c EntriesController) update(ctx *gin.Context) {

}

func (c EntriesController) delete(ctx *gin.Context) {

}
