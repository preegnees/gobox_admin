package storage

import (
	"context"

	models "jwt/pkg/models"
)

func (p *postgres_) SaveUser(ctx context.Context, user models.UserEntity) (err error) {

	query :=
		`
		INSERT INTO users (username, password_hash, email, role) 
		VALUES ($1, $2, $3, $4)
		`

	err = p.client.QueryRow(
		 query, user.Username, user.PasswordHash, user.Email, user.Role,
	).Scan()
	if err != nil {
		return err
	}
	return nil
}
// func (p *postgres_) FindUser(ctx context.Context, username string, password string) (ok bool, err error) {

// 	err = p.client.QueryRow(
// 		ctx, "SELECT exists(SELECT 1 FROM users WHERE )",
// 	).Scan(&user.Username, &user.PasswordHash, &user.Email, &user.Role)
// 	if err != nil {
// 		return err
// 	}
// 	return true, nil
// }

// func (p *postgres_) SaveAppData(ctx context.Context, username string, appTokens []models.Tokens) (err error)
// func (p *postgres_) GetAppData(ctx context.Context, username string) (appTokens []models.Tokens, err error)
