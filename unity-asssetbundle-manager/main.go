package main

import "github.com/gin-gonic/gin"
import "github.com/mosluce/go-playground/unity-asssetbundle-manager/handlers/api"
import sqlite "github.com/mosluce/go-playground/unity-asssetbundle-manager/lib/database/sqlite3"
import "github.com/mosluce/go-playground/unity-asssetbundle-manager/lib/storage"

func main() {
	sqlite.Open("db.sqlite").Migrate()
	storage.Get("storage")

	app := gin.Default()

	api.AssetBundle(app)

	app.Run()
}
