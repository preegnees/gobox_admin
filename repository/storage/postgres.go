package storage

import (
	dto "core/dto"
	// models "core/models"
	"context"
	repository "core/repository"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type postrges_ struct {
	conn *pgx.Conn
}

var _ repository.IStorage = (*postrges_)(nil)

type ConfPostgres struct {
	Cxt      context.Context
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func New(cnf *ConfPostgres) (*postrges_, error){
	if cnf == nil {
		cnf = &ConfPostgres{
			Cxt:      context.TODO(),
			User:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Port:     5432,
			DB:       "postgres",
		}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", (*cnf).User, (*cnf).Password, (*cnf).Host, (*cnf).Port, (*cnf).DB)
	conn, err := pgx.Connect((*cnf).Cxt, url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping((*cnf).Cxt); err != nil {
		return nil, nil
	}

	return &postrges_{
		conn: conn,
	}, nil
}

func (p *postrges_) SaveUser(ctx context.Context, user *dto.DTOUser) (err error) {

	query := 
	`
	INSERT INTO users (username, password) 
	VALUES ($1, $2)
	`
	_, err = p.conn.Exec(ctx, query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil 
}
func (p *postrges_) SaveToken(ctx context.Context, token *dto.DTOTokens) (err error) {
	return nil 
}
func (p *postrges_) RemoveToken(ctx context.Context, token *dto.DTOTokens) (err error) {
	return nil 
}
func (p *postrges_) GetTokens(ctx context.Context, username string) (tokens *[]dto.DTOTokens, err error) {
	return nil, nil 
}
