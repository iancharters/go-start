package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	application "github.com/iancharters/gostart/internal/app"
	"github.com/iancharters/gostart/internal/db"
	"github.com/iancharters/gostart/internal/server"
)

func main() {
	fmt.Println("Starting service...")

	var (
		ctx = context.Background()
		r   = mux.NewRouter()
		cfg = application.NewConfig()
	)

	db, err := db.NewClient(ctx, db.Config{DatabaseURL: cfg.DatabaseURL})
	if err != nil {
		panic("failed to create new db client: " + err.Error())
	}
	defer db.Close()

	var (
		app = application.New(cfg, db)
		svr = server.New(app)
	)

	svr.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
