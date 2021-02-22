package repository

import (
	"context"
	"go-db-simple/internal/comment/entity"
)

// CommentRepository is interface of contrack comment repository
type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindByID(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
