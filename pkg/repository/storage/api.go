package storage

import (
	// "context"
	"database/sql"
	"fmt"
	"os"

	// "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IStorage interface {

}

var _ IStorage = (*postgres_)(nil)

type postgres_ struct {
	client *sql.DB
}

func New() (IStorage, error) {
	
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")

	// url := "user=u1 host=0.0.0.0 port=5432 dbname=d1 sslmode=disable"
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, db)
	
	// url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, db)
	println(url)
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w\n", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	
	return &postgres_{
		client: conn,
	}, nil
}

// // urlExample := "postgres://username:password@localhost:5432/database_name"
// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	os.Exit(1)
// }
// defer conn.Close(context.Background())

// var name string
// var weight int64
// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 	os.Exit(1)
// }

// fmt.Println(name, weight)