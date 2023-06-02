// go:build ignore
package seeder

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const (
	insertQuery = `INSERT INTO users (created_at, updated_at) VALUES `
)

type UserSeeder struct {
	DB *sql.DB
}

func (s *UserSeeder) Seed(ctx context.Context, count int) {
	// Your seeder here
	var stmtBuilder strings.Builder
	stmtBuilder.WriteString(insertQuery)
	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	var users []interface{}
	for i := 0; i < count; i++ {
		user := model.User{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		users = append(users, user.CreatedAt, user.UpdatedAt)

		stmtBuilder.WriteString("(?, ?)")
		if i != count-1 {
			stmtBuilder.WriteString(", ")
		}
	}

	stmt, err := tx.PrepareContext(ctx, stmtBuilder.String())
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, users...)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
}
