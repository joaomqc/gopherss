package ctrl

import (
	"gopherss/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct{}

var articlesRepository = db.ArticlesRepository{}

type Query struct {
	BaseQuery
	Category string `form:"category"`
	Feed     string `form:"feed"`
}

func (c ArticlesController) Register(r *gin.Engine) {
	group := r.Group("/articles")

	group.GET("/", c.getAll)
	group.POST("/", c.create)
	group.PUT("/", c.update)
	group.DELETE("/", c.delete)
}

func (c ArticlesController) getAll(ctx *gin.Context) {
	query := Query{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	articles, err := articlesRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func (c ArticlesController) create(ctx *gin.Context) {

}

func (c ArticlesController) update(ctx *gin.Context) {

}

func (c ArticlesController) delete(ctx *gin.Context) {

}
