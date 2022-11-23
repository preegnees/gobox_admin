package storage

import (
	"context"
	dto "core/dto"
	repository "core/repository"
	models "core/repository/models"
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
	Sslmode  string
}

func New(cnf *ConfPostgres) (*postrges_, error) {
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
	println(url)
	conn, err := pgx.Connect((*cnf).Cxt, url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping((*cnf).Cxt); err != nil {
		return nil, err
	}

	return &postrges_{
		conn: conn,
	}, nil
}

func (p *postrges_) SaveUser(ctx context.Context, user *dto.DTOUser) (err error) {

	query :=
		`
	INSERT INTO users (username, password_hash, user_role, email) 
	VALUES ($1, $2, $3, $4)
	`
	_, err = p.conn.Exec(ctx, query, user.Username, user.PasswordHash, user.UserRole, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *postrges_) FindUserByUsername(ctx context.Context, username string) (*dto.DTOUser, error) {
	query :=
		`
	SELECT user_id, username, password_hash, user_role, email 
	FROM users 
	WHERE username=$1
	`
	user := models.User{}
	err := p.conn.QueryRow(ctx, query, username).Scan(&user.UserId, &user.Username, &user.PasswordHash, &user.UserRole, &user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.DTOUser{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		UserRole:     user.UserRole,
		Email:        user.Email,
	}, nil
}

func (p *postrges_) ChangeRoleUser(ctx context.Context, username string, role string) error {
	return nil
}

func (p *postrges_) ChangeEmailUser(ctx context.Context, username string, role string) error {
	return nil
}

func (p *postrges_) SaveToken(ctx context.Context, token *dto.DTOTokens) error {
	query :=
		`
		SELECT user_id 
		FROM users 
		WHERE username=$1
	`
	user_id := -1
	err := p.conn.QueryRow(ctx, query, token.Username).Scan(&user_id)
	if err != nil {
		return err
	}

	if user_id != -1 {
		query =
			`
		INSERT INTO tokens (temp_token, user_id)
		VALUES ($1, $2)
		`
		_, err = p.conn.Exec(ctx, query, token.Token, user_id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *postrges_) RemoveToken(ctx context.Context, token *dto.DTOTokens) (error) { // сюда нужно передавать DTO STORAGE

	query :=
		`
	DELETE 
	FROM tokens 
	WHERE temp_token=$1
`
	_, err := p.conn.Exec(ctx, query, token.Token)
	if err != nil {
		return err
	}

	return nil
}
func (p *postrges_) GetTokens(ctx context.Context, username string) (tokens *[]dto.DTOTokens, err error) {
	return nil, nil
}
