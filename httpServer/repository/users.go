package repository

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type User struct {
	ID        string `json:"ID"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	// Profile  string `json:"profileImage"`
}

type UserRepo interface {
	GetByID(ctx context.Context, userID string) (*User, error)
	Add(ctx context.Context, u *User, node *snowflake.Node) (string, error)
	Update(ctx context.Context, u *User) error
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(ctx context.Context, userID string) (*User, error) {
	var user User

	err := WithTransaction(us.db, func(tx *sql.Tx) error {
		query := `
			SELECT ID, user_name, first_name, last_name FROM USERS
			WHERE ID = $1
		`
		return tx.QueryRow(query, userID).Scan(
			&user.ID,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
		)
	})

	if err != nil {
		return nil, ErrNotFound
	}

	return &user, nil
}

func (us *UserStore) Add(ctx context.Context, u *User, node *snowflake.Node) (string, error) {
	var id string

	err := WithTransaction(us.db, func(tx *sql.Tx) error {
		nameExist, err := us.isUsernameExist(ctx, u.Username)
		if err != nil {
			return err
		}
		if !nameExist {
			return ErrUsernameExist
		}

		snowflakeID := node.Generate()

		salt, err := generateSalt(16)
		if err != nil {
			return err
		}
		hashedPassword, err := hashPassword(u.Password, salt)
		if err != nil {
			return err
		}

		insertQuery := `
			INSERT INTO USERS (ID, hashed_password, salt, first_name, last_name, user_name)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING ID
		`
		return tx.QueryRowContext(ctx, insertQuery, snowflakeID.String(), hashedPassword, salt, u.Firstname, u.Lastname, u.Username).Scan(&id)
	})

	if err != nil {
		return "", err
	}
	return id, nil
}

func (us *UserStore) Update(ctx context.Context, u *User) error {
	err := WithTransaction(us.db, func(tx *sql.Tx) error {
		updateQuery := `
			UPDATE USERS
			SET first_name = $1, last_name = $2, user_name = $3
			WHERE ID = $4
		`
		res, err := tx.ExecContext(ctx, updateQuery, u.Firstname, u.Lastname, u.Username, u.ID)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ErrNotFound
		}

		return nil
	})

	return err
}

func (us *UserStore) isUsernameExist(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM USERS WHERE user_name = $1)`

	var exists bool
	err := us.db.QueryRowContext(ctx, query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func hashPassword(password, salt string) (string, error) {
	combined := password + salt

	hash := sha256.New()
	_, err := hash.Write([]byte(combined))
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// func main() {
// postgresUserRepo := NewPostgresUserRepo(&sql.DB{})
// 	dynamoDBRepo := NewDynamoDBRepo()
// 	userService := NewUserService(postgresUserRepo)

// 	_ := userService.AddUser(context.Background(), User{})
// }

// func (us *UserStore) Add(ctx context.Context, u *User, node *snowflake.Node) (string, error) {
// 	tx, err := us.db.BeginTx(context.TODO(), nil)
// 	if err != nil {
// 		return "", fmt.Errorf("error starting transaction: %v", err)
// 	}
// 	defer tx.Rollback()

// 	nameExist, err := us.isUsernameExist(ctx, u.Username)
// 	if err != nil {
// 		return "", err
// 	}
// 	if !nameExist {
// 		return "", ErrUsernameExist
// 	}

// 	snowflakeID := node.Generate()

// 	salt, err := generateSalt(16)
// 	if err != nil {
// 		return "", err
// 	}
// 	hashedPassword, err := hashPassword(u.Password, salt)
// 	if err != nil {
// 		return "", err
// 	}

// 	insertQuery := `
//         INSERT INTO USERS (ID, hashed_password, salt, first_name, last_name, user_name)
//         VALUES ($1, $2, $3, $4, $5, $6)
//         RETURNING ID
//     `
// 	var id string
// 	err = tx.QueryRowContext(ctx, insertQuery, snowflakeID.String(), hashedPassword, salt, u.Firstname, u.Lastname, u.Username).Scan(&id)
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return "", fmt.Errorf("error committing transaction: %v", err)
// 	}
// 	return id, nil
// }
