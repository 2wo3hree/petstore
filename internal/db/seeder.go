package db

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedDatabase(pool *pgxpool.Pool) {
	ctx := context.Background()
	seedAuthors(ctx, pool)
	seedBooks(ctx, pool)
	seedUsers(ctx, pool)
}

func seedAuthors(ctx context.Context, pool *pgxpool.Pool) {
	var count int
	err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM authors`).Scan(&count)
	if err != nil {
		log.Fatalf("seedAuthors: count authors: %v", err)
	}
	if count == 0 {
		log.Println("Seeding 10 authors…")
		for i := 0; i < 10; i++ {
			name := gofakeit.Name()
			if _, err := pool.Exec(ctx, `INSERT INTO authors (name) VALUES ($1)`, name); err != nil {
				log.Fatalf("seedAuthors: insert author: %v", err)
			}
		}
	}
}

func seedBooks(ctx context.Context, pool *pgxpool.Pool) {
	var count int
	if err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM books`).Scan(&count); err != nil {
		log.Fatalf("seedBooks: count books: %v", err)
	}
	if count == 0 {
		// забираем все авторы
		rows, err := pool.Query(ctx, `SELECT id FROM authors`)
		if err != nil {
			log.Fatalf("seedBooks: select authors: %v", err)
		}
		defer rows.Close()

		var authorIDs []int
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				log.Fatalf("seedBooks: scan author id: %v", err)
			}
			authorIDs = append(authorIDs, id)
		}
		if len(authorIDs) == 0 {
			log.Fatal("seedBooks: no authors found")
		}

		log.Println("Seeding 100 books…")
		for i := 0; i < 100; i++ {
			title := gofakeit.BookTitle()
			aid := authorIDs[rand.Intn(len(authorIDs))]
			if _, err := pool.Exec(ctx,
				`INSERT INTO books (title, author_id) VALUES ($1, $2)`,
				title, aid,
			); err != nil {
				log.Fatalf("seedBooks: insert book: %v", err)
			}
		}
	}
}

func seedUsers(ctx context.Context, pool *pgxpool.Pool) {
	var count int
	if err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM users`).Scan(&count); err != nil {
		log.Fatalf("seedUsers: count users: %v", err)
	}
	if count < 50 {
		toCreate := 50 - count
		log.Printf("Seeding %d users…\n", toCreate)
		for i := 0; i < toCreate; i++ {
			name := gofakeit.Name()
			if _, err := pool.Exec(ctx, `INSERT INTO users (name) VALUES ($1)`, name); err != nil {
				log.Fatalf("seedUsers: insert user: %v", err)
			}
		}
	}
}
