package repository

import (
	"database/sql"
)

type PostgresUserRepo struct {
	User    UserRepo
	Post    PostRepo
	Comment CommentRepo
	Like    LikeRepo
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{}
}
