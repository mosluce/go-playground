package api

import (
	"net/http"

	"github.com/mosluce/go-playground/unity-asssetbundle-manager/models"

	"github.com/gin-gonic/gin"
	sqlite "github.com/mosluce/go-playground/unity-asssetbundle-manager/lib/database/sqlite3"
	"github.com/mosluce/go-playground/unity-asssetbundle-manager/lib/storage"
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
	ab, err := ctx.FormFile("ab")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	manifest, err := ctx.FormFile("manifest")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	db := sqlite.Open("db.sqlite")
	defer db.Close()

	tx := db.DB.Begin()

	var count int
	tx.Model(&models.AssetBundle{}).Where("Name = ?", ab.Filename).Count(&count)

	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "duplicate assetbundle name",
		})
		tx.Rollback()
		return
	}

	// TODO : 驗證檔案

	// 取得資料夾
	dir := storage.Get("storage")
	// 儲存 assetbundle
	if f := dir.Dir("ab").Save(ab); f.Error != nil {
		ctx.String(http.StatusInternalServerError, f.Error.Error())
		tx.Rollback()
		return
	}
	// 儲存 manifest
	if f := dir.Dir("manifest").Save(manifest); f.Error != nil {
		ctx.String(http.StatusInternalServerError, f.Error.Error())
		tx.Rollback()
		return
	}

	tx.Create(&models.AssetBundle{Name: ab.Filename, Available: true})
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// List of all AssetBundles
func List(ctx *gin.Context) {
	db := sqlite.Open("db.sqlite")
	defer db.Close()

	var abs []models.AssetBundle
	db.DB.Find(&abs)

	ctx.JSON(http.StatusOK, abs)
}

// Get AssetBundle data
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{})
}
