package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fmantinossi/sqlc-app/internal/db"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend - .Net", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "8fa533d4-8297-40a9-834a-c4d32845e8b7",
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend - Go", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "c7229150-fa6b-4311-9bb0-c8b4855b7361 ")
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description)
	}
}
