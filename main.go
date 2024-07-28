package main

import (
	"embed"

	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed db/schema.sql
var schema string

func main() {

	db, err := sql.Open("sqlite3", "~/super-todo/todo.db?mode=rw")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Printf("Migration error : %s\n", err.Error())
	}

	defer db.Close()
	// Create an instance of the app structure
	app := NewApp(db)

	// Create application with options
	err = wails.Run(&options.App{
		Title:         "SuperTodo",
		Width:         500,
		Height:        768,
		DisableResize: true,
		Frameless:     true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			// WindowIsTranslucent:  true,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			// WindowIsTranslucent:  true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
