package postgresql

import (
	"context"
	"github.com/juicyluv/online-store/internal/domain"
)

//language=PostgreSQL
const getUserByIdSQL = `select tel, email, created_at, modified_at from web.users where id = $1`

func (r *repository) GetUserById(ctx context.Context, userId int64) (*domain.User, error) {
	rows, err := r.db.Query(ctx, getUserByIdSQL, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user domain.User

	err = rows.Scan(
		&user.Telephone,
		&user.Email,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		return nil, err
	}

	user.Id = &userId

	return &user, nil
}
