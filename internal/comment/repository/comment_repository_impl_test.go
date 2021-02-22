package repository

import (
	"context"
	"fmt"
	"go-db-simple/internal/comment/entity"
	"testing"

	go_db_simple "go-db-simple/pkg/db"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {

	commentRepository := NewCommentRepository(go_db_simple.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_simple.GetConnection())

	comment, err := commentRepository.FindByID(context.Background(), 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_simple.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
