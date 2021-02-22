package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-db-simple/internal/comment/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

// NewCommentRepository is func for create new comment repository to db
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

// Insert is func of comment repository for implementastion insert comment to db
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlCommand := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, sqlCommand, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.ID = int32(id)
	return comment, nil
}

// FindByID is func of comment repository for implementastion FindByID comment in db
func (repository *commentRepositoryImpl) FindByID(ctx context.Context, id int32) (entity.Comment, error) {
	sqlCommand := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlCommand, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// there are data
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// thare are not data
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

// FindAll is func of comment repository for implementastion FindAll comment in db
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlCommand := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlCommand)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
