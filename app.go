package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	DB "todo/db"
)

// App struct
type App struct {
	ctx     context.Context
	db_ctx  context.Context
	queries DB.Queries
}

// NewApp creates a new App application struct
func NewApp(db *sql.DB) *App {

	queries := DB.New(db)

	db_ctx := context.Background()

	app := &App{db_ctx: db_ctx, queries: *queries}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AddTask(content string) DB.Todo {
	ret, err := a.queries.AddTask(a.db_ctx, content)
	if err != nil {
		log.Println(err.Error())

	}
	return ret
}

func (a *App) GetTasks() []DB.Todo {
	todos, _ := a.queries.GetTodos(a.db_ctx)
	return todos
}

func (a *App) RemoveTask(id int) DB.Todo {
	todo, _ := a.queries.RemoveTask(a.db_ctx, int64(id))
	return todo
}

func (a *App) TickTask(done bool, id int) DB.Todo {
	todo, _ := a.queries.UpdateTask(a.db_ctx, DB.UpdateTaskParams{IsDone: done, ID: int64(id)})
	return todo
}
