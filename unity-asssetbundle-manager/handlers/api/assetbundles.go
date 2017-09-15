package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AssetBundle handles api/assetbundles group
func AssetBundle(app *gin.Engine) {
	g := app.Group("/api/assetbundles")
	{
		g.GET("/", List)
		g.GET("/:name", Get)
		g.POST("/", Create)
	}
}

// Create a new AssetBundle
func Create(ctx *gin.Context) {
	ab, _ := ctx.FormFile("ab")
	manifest, _ := ctx.FormFile("manifest")

	ctx.JSON(http.StatusOK, gin.H{
		"ab":       ab.Filename,
		"manifest": manifest.Filename,
	})
}

// List of all AssetBundles
func List(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{})
}

// Get AssetBundle data
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{})
}
