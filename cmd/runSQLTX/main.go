package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fmantinossi/sqlc-app/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
	CategoryID  string
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("Error on rollback: [%v], original error: [%v]", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseCategory(ctx context.Context, paramsCategory CategoryParams, paramsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          paramsCategory.ID,
			Name:        paramsCategory.Name,
			Description: paramsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          paramsCategory.ID,
			Name:        paramsCourse.Name,
			Description: paramsCourse.Description,
			Price:       paramsCourse.Price,
			CategoryID:  paramsCategory.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	//courseArgs := CourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Go",
	//	Description: sql.NullString{String: "Go Course", Valid: true},
	//	Price:       999.99,
	//}
	//categoryArgs := CategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Go", Valid: true},
	//}
	//
	//courseDB := NewCourseDB(dbConn)
	//err = courseDB.CreateCourseCategory(ctx, categoryArgs, courseArgs)
	//if err != nil {
	//	panic(err)
	//}

}
